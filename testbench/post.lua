-- post.lua - Load testing script for wrk
-- Usage: wrk -t12 -c400 -d30s --script=post.lua http://localhost:3000/process

wrk.method = "POST"
wrk.body   = '{"user_id": 1, "data": "load_test_data_for_performance_comparison"}'
wrk.headers["Content-Type"] = "application/json"

-- Optional: Add randomization
local user_ids = {1, 2, 3, 4, 5}
local data_samples = {
    "performance_test_data_sample_1",
    "performance_test_data_sample_2", 
    "performance_test_data_sample_3",
    "performance_test_data_sample_4",
    "performance_test_data_sample_5"
}

function request()
    local user_id = user_ids[math.random(#user_ids)]
    local data = data_samples[math.random(#data_samples)]
    local body = string.format('{"user_id": %d, "data": "%s"}', user_id, data)
    return wrk.format("POST", nil, nil, body)
end
