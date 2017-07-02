# Homework1 notes
  
## 固定背景不動
background-attachment: fixed;
  
## 讓divs變成inline顯示
display:inline-block;
  
## 背景透明色
background-color: rgba(0, 0, 0, 0.6);
  
## position
欲設定某個div的position時  
先設定其parent div的position為relative(相對)  
position: relative;  
接著把child div的position設定為absolute(絕對)  
position: absolute;  
然後就可以設定child div的top, left, ...等  
top: 14px;  
left: 70px;  
  
## 透過設定margin來置中div
margin: 0 auto;
  
## 設定圓形圖片div
overflow: hidden; /* 內容會被修剪, 並且其餘內容不可見 */  
border-radius: 50%; /* 讓div的外框變成圓形, 最多到50% */  
  
## 設定為3的倍數child element的CSS
<selector>:nth-child(3n) {...}  
  
## 取代inline-block的方法, 可以將container設為flexbox
display: flex; 
align-items: center; /* 垂直置中 */  
justify-content: center /* 水平置中 */  
justify-content: space-around /* 平均分配間隔給content divs */  
flex-wrap: wrap; /* 將超過container邊界的div做wrapping(斷行) */ 
