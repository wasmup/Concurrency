#include <iostream>
#include <string>
#include <thread>
#include <chrono>
#include <mutex>
#include <ctime>
#include <iomanip>
#include <cpprest/http_listener.h> // Install C++ REST SDK: sudo apt install libcpprest-dev

using namespace web::http::experimental::listener;
using namespace web::http;  // Assuming you're using C++ REST SDK for HTTP server

int requestCount = 0;
std::chrono::milliseconds elapsedTime(0);
std::mutex mu;

void printStats() {
    while (true) {
        std::this_thread::sleep_for(std::chrono::seconds(1));
        mu.lock();
        int n = requestCount;
        auto d = elapsedTime;
        mu.unlock();
        std::cout << "Requests: " << n << ", Total Elapsed Time: "
                  << std::setfill('0') << std::setw(2) << d.count() / 1000 << "."
                  << std::setw(3) << d.count() % 1000 << "s\n";
    }
}
#include <iostream>
#include <string>
#include <vector>

std::string base64_encode(const std::string& input) {
    static const char* encoding_table = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
                                        "abcdefghijklmnopqrstuvwxyz"
                                        "0123456789+/";

    std::string encoded;
    encoded.reserve(((input.length() + 2) / 3) * 4); // Pre-allocate space for efficiency

    int i = 0;
    while (i < input.length()) {
        int remaining = input.length() - i;
        int chunk_size = std::min(remaining, 3);

        int value = 0;
        for (int j = 0; j < chunk_size; ++j) {
            value <<= 8;
            value |= input[i + j];
        }

        for (int j = chunk_size; j < 3; ++j) {
            encoded += '=';
        }

        for (int j = 0; j < 4; ++j) {
            int index = (value >> ((3 - j) * 6)) & 0x3F;
            encoded += encoding_table[index];
        }

        i += chunk_size;
    }

    return encoded;
}

void handle_get(http_request request) {
    auto start = std::chrono::system_clock::now();

    std::string q = request.request_uri().query();
    if (q.empty()) {
        q = "Hi";
    }
    // std::string encoded = base64_encode(reinterpret_cast<const unsigned char*>(q.c_str()), q.size());
    request.reply(status_codes::OK, base64_encode(q));

    auto end = std::chrono::system_clock::now();
    auto elapsed = std::chrono::duration_cast<std::chrono::milliseconds>(end - start);

    mu.lock();
    requestCount++;
    elapsedTime += elapsed;
    mu.unlock();
}

int main() {
    std::thread statsThread(printStats);

    http_listener listener("http://localhost:8080");
    listener.support(methods::GET, handle_get);
    try {
        listener.open().wait();
        std::cout << "Listening on http://localhost:8080" << std::endl;
        while(true) {
		    std::this_thread::sleep_for(std::chrono::milliseconds(2000));
	    };
    } catch (const std::exception& ex) {
        std::cerr << "Error: " << ex.what() << std::endl;
    }

    return 0;
}
