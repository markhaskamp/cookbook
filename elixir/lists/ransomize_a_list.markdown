
    def shuffle(list) do
      :random.seed(:erlang.now())
      _shuffle(list, [])
    end
  
    def _shuffle([], acc), do: acc
    def _shuffle(list, acc) do
      next_card_ndx = :random.uniform(Enum.count(list)) - 1
      {leading, [h | t]} = Enum.split(list, next_card_ndx - 1)
      _shuffle(leading ++ t, [h | acc])
    end
