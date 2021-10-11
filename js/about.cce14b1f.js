(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["about"],{9637:function(e,t,a){"use strict";a.r(t);var l=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("recipes-manager")},i=[],n=a("2b0e"),s=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("v-container",[a("v-row",{staticClass:"text-center"},[a("v-col",{attrs:{cols:"12"}},[a("v-data-table",{staticClass:"elevation-1",attrs:{headers:e.headers,items:e.recipes,"items-per-page":5},scopedSlots:e._u([{key:"top",fn:function(){return[a("v-toolbar",{attrs:{flat:""}},[a("v-toolbar-title",[e._v("Recipes CRUD")]),a("v-divider",{staticClass:"mx-4",attrs:{inset:"",vertical:""}}),a("v-spacer"),a("v-dialog",{attrs:{"max-width":"500px"},scopedSlots:e._u([{key:"activator",fn:function(t){var l=t.on,i=t.attrs;return[a("v-btn",e._g(e._b({staticClass:"mb-2"},"v-btn",i,!1),l),[e._v(" New Recipe ")])]}}]),model:{value:e.dialog,callback:function(t){e.dialog=t},expression:"dialog"}},[a("v-card",[a("v-card-title",[a("span",{staticClass:"text-h5"},[e._v("New Recipe")])]),a("v-card-text",[a("v-container",[a("v-row")],1)],1),a("v-card-actions",[a("v-spacer"),a("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:e.close}},[e._v(" Cancel ")]),a("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:e.save}},[e._v(" Save ")])],1)],1)],1),a("v-dialog",{attrs:{"max-width":"500px"},model:{value:e.dialogDelete,callback:function(t){e.dialogDelete=t},expression:"dialogDelete"}},[a("v-card",[a("v-card-title",{staticClass:"text-h5"},[e._v("Are you sure you want to delete this item?")]),a("v-card-actions",[a("v-spacer"),a("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:e.closeDelete}},[e._v("Cancel")]),a("v-btn",{attrs:{color:"blue darken-1",text:""},on:{click:e.deleteItemConfirm}},[e._v("OK")]),a("v-spacer")],1)],1)],1)],1)]},proxy:!0},{key:"item.actions",fn:function(t){var l=t.item;return[a("v-icon",{staticClass:"mr-2",attrs:{small:""},on:{click:function(t){return e.editItem(l)}}},[e._v(" mdi-pencil ")]),a("v-icon",{attrs:{small:""},on:{click:function(t){return e.deleteItem(l)}}},[e._v(" mdi-delete ")])]}}],null,!0)})],1)],1)],1)},o=[],c=n["a"].extend({name:"Recipes",methods:{editItem:function(e){this.dialog=!0},deleteItem:function(e){this.dialogDelete=!0},deleteItemConfirm:function(){this.closeDelete()},close:function(){this.dialog=!1},closeDelete:function(){this.dialogDelete=!1},save:function(){this.close()}},data:function(){return{dialogDelete:!1,dialog:!1,headers:[{text:"Recipe Id",align:"start",value:"id"},{text:"Difficulty",value:"difficulty"},{text:"Description",value:"description"},{text:"Image",value:"image"},{text:"Title",value:"title"},{text:"Subtitle",value:"subtitle"},{text:"Cook time",value:"cook_time"},{text:"Tags",value:"tags"},{text:"Utensils",value:"utensils"},{text:"Allergens",value:"allergens"},{text:"Actions",value:"actions",sortable:!1}],recipes:[{difficulty:"Easy",description:"",image:"",title:"",subtitle:"",cook_time:0,tags:["1"],utensils:["cuchara"],allergens:["a","b"],id:13},{difficulty:"Easy",description:"",image:"",title:"Veamos!",subtitle:"",cook_time:0,tags:["1"],utensils:null,allergens:["a","b"],id:14}]}}}),r=c,d=a("2877"),u=a("6544"),v=a.n(u),m=a("8336"),f=a("b0af"),p=a("99d9"),b=a("62ad"),g=a("a523"),x=a("8fea"),k=a("169a"),_=a("ce7e"),h=a("132d"),C=a("0fd9"),V=a("2fa4"),D=a("71d9"),w=a("2a7f"),y=Object(d["a"])(r,s,o,!1,null,null,null),I=y.exports;v()(y,{VBtn:m["a"],VCard:f["a"],VCardActions:p["a"],VCardText:p["b"],VCardTitle:p["c"],VCol:b["a"],VContainer:g["a"],VDataTable:x["a"],VDialog:k["a"],VDivider:_["a"],VIcon:h["a"],VRow:C["a"],VSpacer:V["a"],VToolbar:D["a"],VToolbarTitle:w["a"]});var R=n["a"].extend({name:"Recipes",components:{RecipesManager:I}}),T=R,S=Object(d["a"])(T,l,i,!1,null,null,null);t["default"]=S.exports}}]);
//# sourceMappingURL=about.cce14b1f.js.map