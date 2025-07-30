const express = require("express");
const mysql = require("mysql2/promise");
const crypto = require("crypto");

class Server {
  constructor() {
    this.app = express();
    this.app.use(express.json());
    this.setupRoutes();
    this.dbPool = null;
  }

  async initDB() {
    // Get database configuration from environment variables
    const dbConfig = {
      host: process.env.DB_HOST || "localhost",
      user: process.env.DB_USER || "root",
      password: process.env.DB_PASSWORD || "abcd",
      database: process.env.DB_NAME || "performance_test",
      waitForConnections: true,
      connectionLimit: 10,
      queueLimit: 0,
      acquireTimeout: 60000,
      timeout: 60000,
    };

    // Create connection pool
    this.dbPool = mysql.createPool(dbConfig);

    try {
      // Test connection with retry logic
      let retries = 5;
      while (retries > 0) {
        try {
          const connection = await this.dbPool.getConnection();
          
          // Create table if it doesn't exist
          await connection.execute(`
            CREATE TABLE IF NOT EXISTS users (
                id INT PRIMARY KEY,
                profile_data VARCHAR(255) NOT NULL
            )
          `);

          // Insert sample data
          const sampleData = [
            { id: 1, profile: "user1_profile_data_with_some_content" },
            { id: 2, profile: "user2_profile_data_with_different_content" },
            { id: 3, profile: "user3_profile_data_with_more_content" },
            { id: 4, profile: "user4_profile_data_with_additional_content" },
            { id: 5, profile: "user5_profile_data_with_extended_content" },
          ];

          for (const data of sampleData) {
            await connection.execute(
              "INSERT IGNORE INTO users (id, profile_data) VALUES (?, ?)",
              [data.id, data.profile]
            );
          }

          connection.release();
          console.log("Database initialized successfully");
          break;
        } catch (error) {
          retries--;
          if (retries === 0) throw error;
          console.log(`Database connection failed, retrying... (${retries} attempts left)`);
          await new Promise(resolve => setTimeout(resolve, 5000));
        }
      }
    } catch (error) {
      console.error("Database initialization failed:", error);
      throw error;
    }
  }

  // CPU-intensive preprocessing: hash the data multiple times
  preprocessData(data) {
    let hash = data;
    // Simulate CPU-intensive work by hashing 1000 times
    for (let i = 0; i < 1000; i++) {
      hash = crypto.createHash("sha256").update(hash).digest("hex");
    }
    return hash.substring(0, 16); // Return first 16 chars
  }

  // Database I/O operation
  async getUserProfile(userID) {
    try {
      const [rows] = await this.dbPool.execute(
        "SELECT profile_data FROM users WHERE id = ?",
        [userID]
      );

      if (rows.length === 0) {
        // Create a mock profile if user doesn't exist
        return `mock_profile_user_${userID}`;
      }

      return rows[0].profile_data;
    } catch (error) {
      throw new Error(`Database error: ${error.message}`);
    }
  }

  // CPU-intensive postprocessing: simulate complex calculations
  postprocessData(preprocessed, profile) {
    const combined = preprocessed + profile;

    // Simulate CPU-intensive calculations
    let result = 0;
    for (let i = 0; i < 100000; i++) {
      result +=
        combined.charCodeAt(i % combined.length) *
        Math.floor(Math.random() * 10);
    }

    // More hashing for additional CPU work
    const finalHash = crypto
      .createHash("sha256")
      .update(combined + result.toString())
      .digest("hex");

    return finalHash.substring(0, 32);
  }

  setupRoutes() {
    // Main processing endpoint
    this.app.post("/process", async (req, res) => {
      const startTime = process.hrtime.bigint();

      try {
        const { user_id: userID, data } = req.body;

        if (!userID || !data) {
          return res.status(400).json({ error: "Missing user_id or data" });
        }

        // Step 1: CPU-bound preprocessing
        const preprocessed = this.preprocessData(data);

        // Step 2: DB I/O
        const profile = await this.getUserProfile(userID);

        // Step 3: CPU-bound postprocessing
        const finalResult = this.postprocessData(preprocessed, profile);

        const endTime = process.hrtime.bigint();
        const processingTime = `${Number(endTime - startTime) / 1000000}ms`;

        const response = {
          processed_data: finalResult,
          user_profile: profile,
          timestamp: new Date().toISOString(),
          processing_time: processingTime,
        };

        res.json(response);
      } catch (error) {
        console.error("Processing error:", error);
        res.status(500).json({ error: "Internal server error" });
      }
    });

    // Health check endpoint
    this.app.get("/health", (req, res) => {
      res.json({ status: "healthy" });
    });
  }

  async start(port = 3000) {
    try {
      await this.initDB();

      this.app.listen(port, "0.0.0.0", () => {
        console.log(`Express server starting on :${port}`);
      });
    } catch (error) {
      console.error("Failed to start server:", error);
      process.exit(1);
    }
  }
}

// Handle graceful shutdown
process.on("SIGINT", () => {
  console.log("Shutting down gracefully...");
  process.exit(0);
});

process.on("SIGTERM", () => {
  console.log("Shutting down gracefully...");
  process.exit(0);
});

// Start the server
const server = new Server();
server.start(3000);
