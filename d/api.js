import express from "express"

const app = express();
app.use(express.json());

const coupons = {Codes:["abc","abcd","abcabc2020"]};

app.get("/coupons", function(req, res){
    console.log(coupons);
    console.log(JSON.stringify(coupons));
    res.send(JSON.stringify(coupons));
})

app.listen("9094");