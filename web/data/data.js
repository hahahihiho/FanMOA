const obj = {
    event : {
        ids : [1,2,3,4,5,6,7],
        names : ["fan1","fan2","fan3","fan4","fan5","fan6","fan7"],
        registrants : ["user_a","user_b","user_c","user_d","user_e","user_f","user_g"],
        celebrities : ["c_1","c_2","c_3","c_4","c_5","c_6","c_7"],
        state : [0,0,1,1,2,1,2],
        close_times : [20220101,20230102,20240205,20211010,20221212,20221212,20221212],
        event_end_time : [20220101,20230102,20240205,20211010,20221212,20221212,20221212],
        max_p : [100,10,200,300,1000,20,100],
        min_p : [10,5,20,30,100,400,500],
        cost : [100,200,1000,3000,5000,6000,5000]
    },

    mypage :{
        mypage_list : ["QR 코드(티켓)","환불","내 참여내역","팬미팅 등록","팬미팅 취소"],
        mypage_link : ["/ticketlist","/refund","/history","/registerevent","/cancellist"].map(l=>"/mypage"+l),
    },
    
    mytickets:[1,3,4],

    makeEntry : function (success,result,error){
        return {success:success, result:result, error: error};
    }
}
    
module.exports = obj;