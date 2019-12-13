# Copyright 2019 Jesse J. Cook
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# hello.feature: hello app features

Feature: Say Hello
    In order to save time on setup code for a problem
    As a developer
    I want a sample hello app written in Go

    Scenario Outline: A GET request is made
        Given the "go-hello" app is running
        When a request is made to: "<path>"
        Then the app should respond: "Hello, you have requested: <path>"
        Examples: Args
            | path   |
            | /      |
            | /World |
