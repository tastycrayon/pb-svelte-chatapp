import{a as n}from"./pocketbase.c6bc3d05.js";import{m as s}from"./stores.71570215.js";import"./ProgressBar.svelte_svelte_type_style_lang.6842832c.js";function c(e){switch(typeof e){case"number":break;case"string":e=+new Date(e);break;default:e=+new Date}const r={timeStyle:"short"},o=new Date,t=new Date(e);if(o.getFullYear()===t.getFullYear()&&o.getMonth()===t.getMonth()&&o.getDate()===t.getDate())return`Today at ${t.toLocaleTimeString([],r)}`;const a=new Date;return a.setDate(a.getDate()-1),a.getFullYear()===t.getFullYear()&&a.getMonth()===t.getMonth()&&a.getDate()===t.getDate()?`Yesterday at ${t.toLocaleTimeString([],r)}`:t.toLocaleDateString()+" "+t.toLocaleTimeString([],r)}const u=(e,r,o)=>r===""||o==""?"/default.jpg":`${n}/api/files/${e}/${r}/${o}`,f=e=>e.toLowerCase().replace(/ /g,"-").replace(/[^\w-]+/g,""),m=()=>new Promise(e=>{const r={type:"confirm",title:"Please Confirm",body:"Are you sure you wish to proceed?",response:o=>{e(o)}};s.trigger(r)});export{f as c,m as f,u as g,c as t};
