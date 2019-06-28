add_library( director_lib
    director.cc
)

target_link_libraries( director_lib PUBLIC
    builder_lib
)

target_include_directories( director_lib PUBLIC
    $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
)