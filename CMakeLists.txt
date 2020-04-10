cmake_minimum_required(VERSION 3.0)

project(Wolfen VERSION 0.1 LANGUAGES C)

set(CMAKE_BUILD_TYPE Debug)
set(CMAKE_C_STANDARD 99)

include_directories(include)

file(GLOB wolfensrc src/*.c)

add_executable(wolfen main.c ${wolfensrc})

find_package(SDL2 REQUIRED)
include_directories(${SDL2_INCLUDE_DIRS})
target_link_libraries(wolfen ${SDL2_LIBRARIES})