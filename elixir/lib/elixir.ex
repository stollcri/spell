defmodule Spell do
  @moduledoc """
  Documentation for Elixir.
  """

  @doc """
  Hello world.

  ## Examples

      iex> Elixir.hello
      :world

  """
  def main(args) do
    args
    |> parse_args
    |> process
  end

  defp parse_args(args) do
    {_, options, _} = OptionParser.parse(args,
      switches: [name: :string]
    )
    options
  end

  def process([]) do
    IO.puts "No arguments given"
  end

  def process(options) do
    for option <- options, do: IO.puts "#{option}"
  end
end
