# @param {String} column_title
# @return {Integer}
def title_to_number(column_title)
  h = [("A".."Z").to_a, (1..26).to_a].transpose.to_h
  return column_title.chars.map { |c| h[c] }.reduce { |r, c| r * 26 + c}
end

p title_to_number("A")
p title_to_number("AA")
p title_to_number("AB")
p title_to_number("ZY")
p title_to_number("FXSHRXW")
