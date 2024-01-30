#include <iostream>
#include <iomanip>
#include <string>
#include <sstream>
#include <codecvt>
#include "httplib.h" // MIT https://github.com/yhirose/cpp-httplib

std::string base64_encode(const std::string& input) {
    // Create a stream to hold the encoded data
    std::stringstream result;

    // Use a codecvt facet for encoding
    std::wstring_convert<std::codecvt_utf8<wchar_t>> converter;
    std::wstring wideInput = converter.from_bytes(input);

    // Iterate through each 3-byte chunk in the input
    for (size_t i = 0; i < wideInput.length(); i += 3) {
        uint32_t chunk = 0;

        // Combine 3 bytes into a 24-bit chunk
        for (size_t j = 0; j < 3; ++j) {
            chunk <<= 8;
            if (i + j < wideInput.length()) {
                chunk |= static_cast<uint8_t>(wideInput[i + j]);
            }
        }

        // Break the 24-bit chunk into 4 6-bit chunks and encode
        for (int j = 18; j >= 0; j -= 6) {
            uint8_t index = static_cast<uint8_t>((chunk >> j) & 0x3F);
            result << "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"[index];
        }
    }

    // Pad with '=' characters if necessary
    size_t padding = input.length() % 3;
    for (size_t i = 0; i < padding; ++i) {
        result << '=';
    }

    return result.str();
}

void handler(const httplib::Request& req, httplib::Response& res) {
    std::string q = req.get_param_value("q");
    if (q.empty()) {
        q = "Hi";
    }

    std::string encoded = base64_encode(  q );
    res.set_content(encoded, "text/plain");
}

int main() {
    httplib::Server svr;

    svr.Get("/", [](const httplib::Request& /*req*/, httplib::Response& res) {
        res.status = 200;
    });

    // svr.Get("/", handler); // Requests/sec:    295.26

    svr.listen("localhost", 8080);

    return 0;
}
