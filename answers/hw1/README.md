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
overflow: hidden; (內容會被修剪, 並且其餘內容不可見)  
border-radius: 50%; (讓div的外框變成圓形, 最多到50%)  
  
## 設定為3的倍數child element的CSS
<selector>:nth-child(3n) {...}  
  
## 取代inline-block的方法, 可以將container設為flexbox
display: flex;  
align-items: center; (垂直置中)  
justify-content: center; (水平置中)  
justify-content: space-around; (平均分配間隔給content divs)   
flex-wrap: wrap; (將超過container邊界的div做wrapping)  

## 讓文字太長出現...的做法
white-space: nowrap; (文字死都不換行)   
text-overflow: ellipsis; (出現點點點的樣式)   
overflow: hidden; (把多的裁掉)   
width:200px; (設定範圍)   

## 參考文獻
[CSS 語法，文字換行，強迫不換行。](https://www.puritys.me/docs-blog/article-31-CSS-%E8%AA%9E%E6%B3%95%EF%BC%8C%E6%96%87%E5%AD%97%E6%8F%9B%E8%A1%8C%EF%BC%8C%E5%BC%B7%E8%BF%AB%E4%B8%8D%E6%8F%9B%E8%A1%8C%E3%80%82.html)

