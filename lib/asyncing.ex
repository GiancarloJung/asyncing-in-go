defmodule Asyncing do
  @moduledoc """
  Documentation for `Asyncing`.
  """

  @doc """
  Hello world.

  ## Examples

      iex> Asyncing.hello()
      :world

  """
  def hello do
    :world
  end

  def run do
    receive do
      msg when is_binary(msg) -> IO.puts(msg)
      _ -> IO.puts("Hello World")
    end
  end
end
