-module(module1).
-export([print_message/0, use_dependency/0]).

print_message() ->
    io:format("Message from Module 1!~n").

use_dependency() ->
    module2:print_message().


