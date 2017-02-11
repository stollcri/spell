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
    options = args |> parse_args
    wordlist = get_wordlist(System.get_env("SPELL_FILE"))
    process(options, wordlist)
  end

  defp parse_args(args) do
    {switches, options, _} = OptionParser.parse(args,
      switches: [name: :string]
    )
    options
  end

  defp get_wordlist(nil) do
    File.read!(".spell")
  end
  defp get_wordlist(wordlist_file) do
    File.read!(wordlist_file)
  end

  def process([], _) do end
  def process(options, wordlist) do
    for option <- options, do: Dp.bestmatch(option, wordlist)
  end
end
