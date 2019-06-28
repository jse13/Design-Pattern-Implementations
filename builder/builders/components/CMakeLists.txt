add_library( component_lib
    component_A.cc
    component_B.cc
)

target_include_directories( component_lib PUBLIC
    $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
)