cmake_minimum_required( VERSION 3.1 )

project( BUILDER )

set ( CMAKE_CXX_STANDARD 17 )

add_subdirectory( "${PROJECT_SOURCE_DIR}/builders" )
add_subdirectory( "${PROJECT_SOURCE_DIR}/director" )

add_executable( builder main.cc )

target_link_libraries( builder director_lib )