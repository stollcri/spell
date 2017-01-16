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
    args |> parse_args |> process
  end

  def process([]) do
    IO.puts "No arguments given"
  end

  def process(options) do
    # IO.puts "Hello #{options[:name]}"
    IO.puts "Hello #{options[1]}"
  end

  defp parse_args(args) do
    # {options, _, _} = OptionParser.parse(args,
    #   switches: [name: :string]
    # )
    {_, options, _} = OptionParser.parse(args,
      switches: [name: :string]
    )
    options
  end
end
