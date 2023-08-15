


def braces():
    stack = []


    A = '{{}}}{}}'
    #A = '{}{}{}{}'
    ans = 0
    counter = 0


    #
    #4


    for i in A:
        
        if i == '{':
            stack.append(i)

        else:
            if len(stack) != 0:
                top = stack[-1]
                # print(top)
                
                if top == '{':
                    stack.pop()
                    counter += 2
                    # print(counter)
                else:
                    print("he")
                    ans = max(ans, counter)    
                    
                    counter = 0
                
                
                
                    
    ans = max(ans, counter)     
             
    
    return ans
    
print(braces())
    