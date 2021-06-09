# dbt-cmake-test
A sample project on how to make cmake understand dbt projects
## Usage
Build with dbt: ` dbt sync && dbt build //...`

Or build with cmake: `mkdir build; cd build; cmake .. && make`

They both work!
