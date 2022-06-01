defmodule AsyncingTest do
  use ExUnit.Case
  doctest Asyncing

  test "greets the world" do
    assert Asyncing.hello() == :world
  end

  test "greets from another process" do
    pid = spawn(fn() -> Asyncing.hello() end)

    assert true == Process.alive?(pid)
  end

  test "greets from with custom message" do
    pid = spawn(fn() -> Asyncing.run() end)

    pid |> send("Hello Go")

    assert true == Process.alive?(pid)
  end
end
