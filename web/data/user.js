
const fs = require('fs')
// const DB = require("./idpw.json")

const DB_PATH = "./idpw.json"
const db_utils = {
    DB : require(DB_PATH),
    read : () => require(DB_PATH),
    write : (id,pw) => {
        const db = db_utils.read();
        db[id] = pw;
        db_utils.save(db);
    },
    save : (db) => {
        try {
            fs.writeFileSync(DB_PATH, db)
            return true;
        } catch (err) {
            console.error(err);
            return false;
        }
    },
    login : (id,pw) ={

    }
}

const check_utils = {
    isIn : (id)=>{
        if(id in Object.keys(db_utils.read())){
            return true
        } else {
            return false
        }
    },
    isCorrect : (id,pw) => {
        if(check_utils.isIn(id) && db_util.DB[id] == pw){
            return true
        } else {
            return false
        }
    }
}

const obj = {
    signup : (id,pw) => {
        if(check_utils.isIn(id)){ return false; }
        else {
            db_utils.write(id,pw); 
            return true; 
        }
    },
    login : (id,pw) => {
        if(check_utils.isCorrect(id,pw)){ return true; }
        else { return false; }
    }

}
    
exports = module.exports = obj