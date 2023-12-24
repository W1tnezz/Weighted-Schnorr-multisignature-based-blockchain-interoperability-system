
import faulthandler;
faulthandler.enable()

import matplotlib

matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib import font_manager
import numpy as np

fig, ax = plt.subplots()
font = font_manager.FontProperties(fname="/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc")
sig_nums = [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50]


bls_gas_cost = []

ecdsa_gas_cost = [111780, 133906, 133906, 156312, 156312, 179064, 179064, 202067, 202067,225379, 225379, 248947, 248947, 272892, 272892, 297081, 297081, 321545,321545, 346234, 346234, 371355, 371355, 396713, 396713, 422363, 422363,448210, 448210, 474523, 474523, 501050, 501050, 527852, 527852, 554827,554827, 582333, 582333, 610029, 610029, 638017, 638017, 666154, 666154,694870, 694870, 723735]

schnorr_gas_cost = [75191, 85824, 85824, 96469, 96469, 107126, 107126, 117771, 117771, 128404, 128404, 139061, 139061, 149694, 149694, 160327, 160327, 170972, 170972, 181629, 181629, 192274, 192274, 202919, 202919, 213576, 213576, 224209, 224209, 234854, 234854, 245487, 245487, 256144, 256144, 266777, 266777, 277410, 277410, 288043, 288043, 298676, 298676, 309309, 309309, 319966, 319966, 330599, 330599, 341256, 341256, 351901, 362546, 373179, 383812, 394469, 405102, 415759, 426416, 437073, 447718, 458351, 468996, 479641, 490298, 500931, 511576, 522221, 532878, 543523, 554156, 564813, 575446]
plt.plot(node_nums, bls_gas_cost[0:48],  color='deepskyblue', label="BLS-threshold")
plt.plot(node_nums, ecdsa_gas_cost[0:48],  color='orange', label="ECDSA")
plt.plot(node_nums[0:14], onchain_gas_cost[0:14],  color='r', label="on-chain")
plt.plot(node_nums, schnorr_gas_cost[0:48], color='green', label="Schnorr-multi")
plt.gcf().subplots_adjust(left=0.15,top=0.9,bottom=0.1)
plt.xlabel("Number of oracle nodes", fontproperties=font)  # 横坐标名字
plt.ylabel("Gas consumption", fontproperties=font)  # 纵坐标名字
plt.legend()
my_x_ticks = np.arange(0, 51, 5)
my_y_ticks = np.arange(0, 750001, 50000)
plt.xticks(my_x_ticks)
plt.yticks(my_y_ticks)
fig.savefig('gas消耗对比.svg', dpi=3200, format='svg')


