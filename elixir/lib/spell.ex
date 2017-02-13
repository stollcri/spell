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
    {_, options, _} = OptionParser.parse(args,
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
  def process(words, wordlist) do
    for word <- words, do: find_match(word, String.split(wordlist, "\n"))
  end

  def find_match(word, wordlist) do
    match = Dp.best_match(word, wordlist)
    IO.puts "MATCH: #{match}"
  end
end
