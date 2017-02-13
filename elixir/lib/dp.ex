defmodule Dp do
  @moduledoc """
  Documentation for Elixir.
  """

  @doc """
  Hello world.

  ## Examples

      iex> Elixir.hello
      :world

  """
  def score(_, "") do end
  def score(word, possibility) do
    String.length(possibility)
  end

  def score_all(word, possibilities) do
    for possibility <- possibilities, do: {possibility, score(word, possibility)}
  end

  def best_match(word, possibilities) do
    word_scores = score_all(word, possibilities)

    {maxword, _} = Enum.reduce word_scores, {"", 0}, fn({word, score}, {accword, acc}) ->
      if word != "" and score >= acc do
        IO.puts "A #{word} #{score} (#{accword})"
        {word, score}
      else
        IO.puts "B #{word} #{acc} (#{accword})"
        {accword, acc}
      end
    end
    maxword
  end
end
