
# def make_kata(input_string):
#     """Test finished in 2.167048931121826 seconds """
#     input_string = input_string.lower()
#     output = ''
#     inputted_list = list(input_string)
#     inputted_set = set(inputted_list)
#     for symbol in inputted_set:
#         inputted_list.remove(symbol)
#     for symbol in input_string:
#         output += ')' if symbol in inputted_list else '(' 
#     return output

def make_kata(inputted):
    """ Test finished in 0.9294266700744629 seconds """
    d = {}
    inputted = inputted.lower()
    outputted = ''
    for symbol in inputted:
        d[symbol] = ')' if symbol in d else '('
    for symbol in inputted:
        outputted += d[symbol]
    return outputted

# def make_kata(inputted):
#     """ Test finished in 0.9294266700744629 seconds """
#     d = {}
#     def f(stri):
#         length = len(stri)
#         if length == 2:
#             stri = stri.lover()
#             if 
#             if stri[0] == stri[1]:
#                 d[stri[0]] = ')'

#     d = set()
#     outputted = ''
#     outputted_list = []
#     for symbol in inputted:
#         symbol = symbol.lower()
#         outputted_list.append() = ')' if symbol in d else '('
#     for symbol in inputted:
#         outputted += d[symbol]
#     return outputted

