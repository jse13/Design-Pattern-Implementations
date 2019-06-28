add_library( builder_lib
    builder_A.cc
    builder_B.cc
)

add_subdirectory( "${CMAKE_CURRENT_SOURCE_DIR}/components" )

target_link_libraries( builder_lib PUBLIC
    component_lib
)

target_include_directories( builder_lib PUBLIC
    $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
)