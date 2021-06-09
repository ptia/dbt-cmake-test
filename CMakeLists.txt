cmake_minimum_required(VERSION 3.8)

project(dbt-test)

#################################### DBT COMPATIBILITY ###################################
# Includes in this project are meant for dbt, so they are relative to dbt's DEPS folder,
# where all modules, including the main project, would have their own folder.
# So every include to files in this project is prefixed by the name of the project folder
# (dbt-cmake-test here).
# For example, to include pkg/header.h we use #include "dbt-cmake-test/pkg/header.h"
# 
# To make CMake understand this structure, we recreate the DEPS folder inside CMake's 
# ${CMAKE_BINARY_DIR} folder. This is NOT in the same place as dbt's DEPS or BUILD folder,
# so CMake and dbt are kept separate, and can be used at the same time.

# Make the CMAKE_DEPS folder (full path ${DBT_INCLUDE_ROOT}) where we recreate dbt's DEPS
set(DBT_INCLUDE_ROOT ${CMAKE_BINARY_DIR}/CMAKE_DEPS)
file(MAKE_DIRECTORY ${DBT_INCLUDE_ROOT})

# This project's dbt name (its main directory name, not necessarily CMake's PROJECT_NAME)
get_filename_component(
    DBT_PROJECT_NAME 
    ${CMAKE_SOURCE_DIR} NAME
    )

# Symlink this project's source dir in CMAKE_DEPS (${DBT_INCLUDE_ROOT})
execute_process(
    COMMAND ln -sfn ${CMAKE_SOURCE_DIR} ${DBT_INCLUDE_ROOT}/${DBT_PROJECT_NAME}
    )

# Copy or symlink any other dependencies in the same ${DBT_INCLUDE_ROOT} folder
# (there aren't any here)

# Make includes relative to the CMAKE_DEPS folder
include_directories(${DBT_INCLUDE_ROOT})
#########################################################################################


add_subdirectory(pkg)
