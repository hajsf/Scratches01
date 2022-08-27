span = test.loc[(test['DATE1'] >= '2022-04-15') & (test['DATE1'] < '2022-05-15')]

fig, ax = plt.subplots(figsize=(25, 5))

plt.plot(span['DATE1'], span['QTY'], label='Historical Data')
plt.plot(span['DATE1'], span['prediction'], marker='o', linestyle='--', color='r', label='Predicted Data')
plt.ylabel('Qty')
plt.legend(['Historical Data', 'Predicted Data'])
plt.show()




