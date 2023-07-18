-module(module2).
-export([print_message/0, use_dependency/0]).

print_message() ->
    io:format("Message from Module 2!~n").

use_dependency() ->
    module1:print_message().
