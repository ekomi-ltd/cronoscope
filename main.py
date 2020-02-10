#!/usr/bin/python
import sys
import time

def get_list_item(index, cast, default):
    try:
        return cast(sys.argv[index])
    except IndexError, ValueError:
        return default

def get_limit():
    return get_list_item(1, int, 10)
    
def get_delay():
    return get_list_item(2, float, 0.2)

def main():
    limit = get_limit()
    delay = get_delay()

    for i in range(0, limit):
        print("Step %02d/%d: sleeping for %0.2f seconds..." % (i + 1, limit, delay))
        sys.stdout.flush()
        time.sleep(delay)

if __name__ == "__main__":
    main()