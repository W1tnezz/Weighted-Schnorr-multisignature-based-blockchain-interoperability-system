import random
node_nums = [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50]

bls_gas_cost = [271943, 242919, 256943, 270943, 256931, 242943, 242919, 242943, 270943, 256943, 284943, 256943, 256943, 284943, 242943, 242943, 242943, 242943,242943, 340943, 298943, 256943, 256943, 242931, 270943, 242943, 256943,270931, 256943, 242943, 270943, 242943, 242943, 242943, 256943, 256943,326943, 242943, 242931, 242907, 256943, 256931, 242943, 242943, 270931,256931, 242943, 242943]

ecdsa_gas_cost = [111780, 133906, 133906, 156312, 156312, 179064, 179064, 202067, 202067,225379, 225379, 248947, 248947, 272892, 272892, 297081, 297081, 321545,321545, 346234, 346234, 371355, 371355, 396713, 396713, 422363, 422363,448210, 448210, 474523, 474523, 501050, 501050, 527852, 527852, 554827,554827, 582333, 582333, 610029, 610029, 638017, 638017, 666154, 666154,694870, 694870, 723735]

onchain_gas_cost = [206620,  280465,  280465,  354310,  354310,  428155,  428155,  502000,502000,  575845,  575845,  649690,  649690,  723535,  723535,  797380,797380,  871225,  871225,  945070,  945070, 1018915, 1018915, 1092760,1092760, 1166605, 1166605, 1240450, 1240450, 1314295, 1314295, 1388140,1388140, 1461985, 1461985, 1535830, 1535830, 1609675, 1609675, 1683520, 1683520, 1757365, 1757365, 1831210, 1831210, 1905055, 1905055, 1978900]

schnorr_gas_cost_temp = [75191, 96457, 117771]

schnorr_gas_cost = [75191, 85824, 96469, 107126, 117771, 128404, 139061, 149694, 160327, 170972, 181629, 192274, 202919, 213576, 224209, 234854, 245487, 256144, 266777, 277410, 288043, 298676, 309309, 319966, 330599, 341256, 351901, 362546, 373179, 383812, 394469, 405102, 415759, 426416, 437073, 447718, 458351, 468996, 479641, 490298, 500931, 511576, 522221, 532878, 543523, 554156, 564813, 575446]

transfer_cost = [10645, 10633, 10657]

bls_verify_cost = 113969

schnorr_verify_cost = 15923

ecdsa_verify_cost = 4359

ecdsa_verify_costs = [8718, 13077, 17436, 21795, 26154, 30513, 34872, 39231, 43590, 47949, 52308, 56667, 61026, 65385, 69744, 74103, 78462, 82821, 87180, 91539, 95898, 100257, 104616, 108975, 113334, 117693, 122052, 126411, 130770, 135129, 139488, 143847, 148206, 152565, 156924, 161283, 165642, 170001, 174360, 178719, 183078, 187437, 191796, 196155, 200514, 204873, 209232, 213591]

bls_verify_costs = [113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969, 113969]

schnorr_verify_costs = [15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923, 15923]

print(len(bls_verify_costs))