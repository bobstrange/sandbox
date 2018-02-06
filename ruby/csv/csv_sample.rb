require 'csv'

filename = 'sample.csv'

# やりたいこと
# CSVを読み込んで
# 不要なデータ 2018-01-26 以前のデータをフィルタして
# request_time順に並べる
table = CSV.table(filename)

puts table.select { |d|
  Date.strptime(d[:day], '%Y-%m-%d') > Date.new(2018, 1, 26)
}.sort_by { |d|
  -d[:request_time]
}
