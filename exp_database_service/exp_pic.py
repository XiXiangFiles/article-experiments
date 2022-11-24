import matplotlib.pyplot as plt


base = 150000
da = []
ta = []
ta2 = []

workers = [1, 3, 6 ,12 ,24, 48]

i = 0
for e in [34.9,103.9,219.6,432.8,1006.6,1739]:
    da.append((e/workers[i])/base)
    i += 1 

j = 0
for e in [64.8,239.5,410.3,882,2397,6110]:
    ta.append((e/workers[j])/base)
    j += 1 


k = 0
for e in [72,133,256,665,1625,3758]:
    ta2.append((e/workers[k])/base)
    k += 1 


fig, ax = plt.subplots()
ax.plot( workers, da, 'o-', label='directly access')
ax.plot( workers, ta,'+-', label='through gRPC access')
ax.plot( workers, ta2,'x-', label='through gRPC access with 2 replica')


ax.legend(loc='best', shadow=True)
ax.set_xlabel('workers')
ax.set_ylabel('secs')
ax.set_title('A Single Query Time Experiment')
plt.show()