
def make_kata(input_string):
    input_string = input_string.lower()
    output = ''
    inputted_list = list(input_string)
    inputted_set = set(inputted_list)
    for symbol in inputted_set:
        inputted_list.remove(symbol)
    for symbol in input_string:
        output += ')' if symbol in inputted_list else '(' 
    return output
