import heapq

def get_median(min_heap, max_heap):
    if len(min_heap) > len(max_heap):
        min_val = heapq.heappop(min_heap)
        heapq.heappush(min_heap, min_val)
        return min_val
    elif len(min_heap) < len(max_heap):
        max_val = heapq.heappop(max_heap)
        heapq.heappush(max_heap, max_val)
        return max_val
    else:
        min_val = heapq.heappop(min_heap)
        heapq.heappush(min_heap, min_val)
        max_val = heapq.heappop(max_heap)
        heapq.heappush(max_heap, max_val)
        return (min_val + max_val) / 2
    
def add(num, min_heap, max_heap):
    if len(min_heap) + len(max_heap) <= 1:
        heapq.heappush(max_heap, num)
        return
    
    median = get_median(min_heap, max_heap)
    if num > median:
        heapq.heappush(min_heap, num)
    else: 
        heapq.heappush(max_heap, num)

def rebalance(min_heap, max_heap):
    if len(min_heap) > len(max_heap) + 1:
        root = heapq.heappop(min_heap)
        heapq.heappush(max_heap, root)
    elif len(max_heap) > len(min_heap) + 1:
        root = heapq.heappop(max_heap)
        heapq.heappush(min_heap, root)

def print_median(min_heap, max_heap):
    print(get_median(min_heap, max_heap))

def running_median(stream):
    min_heap = []
    max_heap = []
    for num in stream: 
        add(num, min_heap, max_heap)
        rebalance(min_heap, max_heap)
        print_median(min_heap, max_heap)


running_median([1, 2, 3])