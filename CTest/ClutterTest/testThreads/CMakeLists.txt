cmake_minimum_required(VERSION 3.6)
project(testThreads)

# test support for C++11
include(CheckCXXCompilerFlag)
CHECK_CXX_COMPILER_FLAG("-std=c++11" COMPILER_SUPPORTS_CXX11)
CHECK_CXX_COMPILER_FLAG("-std=c++0x" COMPILER_SUPPORTS_CXX0X)

if(COMPILER_SUPPORTS_CXX11)
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11")
elseif(COMPILER_SUPPORTS_CXX0X)
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++0x")
else()
    message(STATUS "The compiler ${CMAKE_CXX_COMPILER} has no C++11 support.")
    message("Please use a different C++ compiler.")
endif()
set(SOURCE_FILES
    main.c)

add_executable(testThreads ${SOURCE_FILES})
# add clutter support
# pkg name can run follow command to check
# pkg-config --list-all | grep YOU_WANT_CHECK
find_package(PkgConfig REQUIRED)
pkg_check_modules(CLUTTER REQUIRED clutter-1.0)
pkg_check_modules(GLIB2 REQUIRED glib-2.0)

include_directories(${CLUTTER_INCLUDE_DIRS} ${GLIB2_INCLUDE_DIRS})

link_directories(${CLUTTER_LIBRARY_DIRS} ${GLIB2_INCLUDE_DIRS})

add_definitions(${CLUTTER_CFLAGS_OTHER} ${GLIB2_CFLAGS_OTHER})

target_link_libraries(${PROJECT_NAME} ${CLUTTER_LIBRARIES} ${GLIB2_LIBRARIES})
