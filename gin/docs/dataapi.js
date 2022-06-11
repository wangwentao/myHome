
// master page
function getUserScore() {
  return userscore;
}

function getWaterData(monthkey) {

  let rest = userwaterdata[monthkey]
  if(!rest) {
    rest = {}
  }
  return rest
}

function getAnalysisData(type, month) {

  return resultData
}

const userscore = {
  userId: 10,
  userName: "王子安",
  userScore: 260
}


const userwaterdata = {
  "2022-03": {
    month: "2022-03",
    income: +234.00,
    outcome:-10.33,
    days: [
      {
        day:"2022-03-21",
        income: +234.00,
        outcome:-10.33,
        list:[
          {
            recordid:1341235225,
            date: "2022-03-21",
            type: 1,      
            name: "奖金",
            score: 123.00,
            user: "",
            remark: "备注",
            typeimg: "jiangjin.png"
          },
          {
            recordid: 1341235225,
            date: "2022-03-21",
            type: 2,//记账类型          
            name: "购物消费",
            score: -123.00,//花费金额
            user: "",
            remark: "备注2",
            typeimg: "gouwu.png"
          }
        ],
      }     
    ]
  },
  "2022-04": {
    month: "2022-04",
    income: +234.00,
    outcome:-10.33,
    days: [
      {
        day:"2022-04-21",
        income: +234.00,
        outcome:-10.33,
        list:[
          {
            recordid:1341235225,
            date: "2022-04-21",
            type: 1,      
            name: "奖金",
            score: 123.00,
            user: "",
            remark: "备注",
            typeimg: "jiangjin.png"
          },
          {
            recordid: 1341235225,
            date: "2022-04-21",
            type: 2,//记账类型          
            name: "购物消费",
            score: -123.00,//花费金额
            user: "",
            remark: "备注2",
            typeimg: "gouwu.png"
          }
        ],
      }     
    ]

  },
  "2022-05": {
    month: "2022-05",
    income: +234.00,
    outcome:-10.33,
    days: [
      {
        day:"2022-05-21",
        income: +234.00,
        outcome:-10.33,
        list:[
          {
            recordid:1341235225,
            date: "2022-05-21",
            type: 1,      
            name: "奖金",
            score: 123.00,
            user: "",
            remark: "备注",
            typeimg: "jiangjin.png"
          },
          {
            recordid: 1341235225,
            date: "2022-05-21",
            type: 2,//记账类型          
            name: "购物消费",
            score: -123.00,//花费金额
            user: "",
            remark: "备注2",
            typeimg: "gouwu.png"
          }
        ],
      }     
    ]
  },
  "2022-06": {
    month: "2022-05",
    income: +234.00,
    outcome:-10.33,
    days: [
      {
        day:"2022-06-21",
        income: +234.00,
        outcome:-10.33,
        list:[
          {
            recordid:1341235225,
            date: "2022-05-21",
            type: 1,      
            name: "奖金",
            score: 123.00,
            user: "",
            remark: "备注",
            typeimg: "jiangjin.png"
          },
          {
            recordid: 1341235225,
            date: "2022-06-21",
            type: 2,//记账类型          
            name: "购物消费",
            score: -123.00,//花费金额
            user: "",
            remark: "备注2",
            typeimg: "gouwu.png"
          }
        ],
      }     
    ]
  }
}

// score page

function getScoreItmes() {

  return scoreItems;
}

const scoreItems = {
  defaultItem: {
    "id": "wczy",
    "type": 1,
    "name": "完成作业",
    "score": 2,
    "image": "jianzhi.png"
  },
  incomeItems : [{
    "id": "wczy",
    "type": 1,
    "name": "完成作业",
    "score": 2,
    "image": "jianzhi.png"
  },
  {
    "id": "lsby",
    "type": 2,
    "name": "老师表扬",
    "score": 1,
    "image": "gongzi.png"
  },
  {
    "id": "zl",
    "type": 3,
    "name": "自爱自律",
    "score": 3,
    "image": "hongbao.png"
  },
  {
    "id": "jc",
    "type": 4,
    "name": "坚持顽强",
    "score": 3,
    "image": "touzi.png"
  },
  {
    "id": "cf",
    "type": 5,
    "name": "吃饭",
    "score": 1,
    "image": "jiangjin.png"
  },
  {
    "id": "cdjw",
    "type": 6,
    "name": "承担家务",
    "score": 2,
    "image": "butie.png"
  },
  {
    "id": "wjzl",
    "type": 7,
    "name": "整理玩具",
    "score": 1,
    "image": "lijin.png"
  },
  {
    "id": "more",
    "type": 8,
    "name": "更多",
    "image": "qita.png"
  }],
  expendItems : [{
    "id": 'szh',
    "type": 10,
    "name": "说脏话",
    "score": -2,
    "image": "canyin.png"
  },
  {
    "id": 'sjh',
    "type": 11,
    "name": "说假话",
    "score": -4,
    "image": "jiaotong.png"
  },
  {
    "id": 'wgdr',
    "type": 12,
    "name": "无故打人",
    "score": -3,
    "image": "jujia.png"
  },
  {
    "id": 'lspp',
    "type": 13,
    "name": "老师批评",
    "score": -2,
    "image": "gouwu.png"
  }, {
    "id": 'lrwj',
    "type": 14,
    "name": "乱扔玩具",
    "score": -2,
    "image": "tongxun.png"
  }, {
    "id": 'wldz',
    "type": 15,
    "name": "无理顶嘴",
    "score": -2,
    "image": "xuexi.png"
  }, {
    "id": 'lfls',
    "type": 16,
    "name": "浪费粮食",
    "score": -2,
    "image": "jiankang.png"
  }, {
    "id": 'sf',
    "type": 17,
    "name": "剩饭",
    "score": -1,
    "image": "yule.png"
  },{
    "id": "more",
    "type": 18,
    "name": "更多",
    "image": "qita.png"
  }]
}

const resultData = {
  monthScore: 126,
  scoreSeries: [{
    name: '成交量1',
    data: 15,
    stroke: false
  }, {
    name: '成交量2',
    data: 35,
    stroke: false
  }, {
    name: '成交量3',
    data: 78,
    stroke: false
  }, {
    name: '成交量4',
    data: 63,
    stroke: false
  }, {
    name: '成交量4',
    data: 63,
    stroke: false
  }, {
    name: '成交量4',
    data: 63,
    stroke: false
  }, {
    name: '成交量4',
    data: 63,
    stroke: false
  }]
}




module.exports = {
  getUserScore: getUserScore,
  getWaterData: getWaterData,
  getScoreItmes: getScoreItmes,
  getAnalysisData: getAnalysisData
}
