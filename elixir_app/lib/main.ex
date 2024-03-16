# Load into iex with c("./lib/main.ex", "")

defmodule Main do
    def hello do
        :world
    end
end

defmodule NumberAdder do
    def run do
        {amount_of_random_numbers, number_of_groups} = get_input()
        random_number_list = Enum.map(1..amount_of_random_numbers, fn _ -> :rand.uniform(100) end)
        split_list = Enum.chunk_every(number_list, div(Enum.count(random_number_list), number_of_groups))
        pids = Enum.map(split_list, &spawn_link(NumberAdder, :begin_summing, [&1]))
        IO.inspect(pids)
    end

    defp get_input do
        IO.puts("Please enter the amount of random numbers you want to generate: ")
        amount_of_nums = IO.gets("") |> String.trim()
        amount_of_groups = String.length(amount_of_nums)
        random_nums =
            case Integer.parse(amount_of_nums) do
                {n, _} when is_number(n) -> n
                _ ->
                    IO.puts "Invalid input. Defaulting to 100."
                    100
            end
        {random_nums, amount_of_groups}
    end

    def begin_summing(nums) do
        sum = Enum.reduce(nums, 0, fn x, acc -> x + acc end)
        send(self(), {self(), sum})
    end
end

NumberAdder.run()
