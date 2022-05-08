!function(e,t){if("object"==typeof exports&&"object"==typeof module)module.exports=t();else if("function"==typeof define&&define.amd)define([],t);else{var r=t();for(var n in r)("object"==typeof exports?exports:e)[n]=r[n]}}(self,(()=>(()=>{"use strict";var e={422:(e,t)=>{Object.defineProperty(t,"__esModule",{value:!0}),t.ValidateDisplayName=t.ValidateUserName=void 0,t.ValidateUserName=function(e){return e.toLowerCase().replace(/\s/g,"")},t.ValidateDisplayName=function(e){return e.toLowerCase().split(".")[0]}},85:function(e,t,r){var n=this&&this.__awaiter||function(e,t,r,n){return new(r||(r=Promise))((function(o,i){function s(e){try{c(n.next(e))}catch(e){i(e)}}function a(e){try{c(n.throw(e))}catch(e){i(e)}}function c(e){var t;e.done?o(e.value):(t=e.value,t instanceof r?t:new r((function(e){e(t)}))).then(s,a)}c((n=n.apply(e,t||[])).next())}))};Object.defineProperty(t,"__esModule",{value:!0}),t.startUserLogin=void 0;const o=r(825),i=r(454),s=r(422),a=r(441);t.startUserLogin=function(e,t=""){return n(this,void 0,void 0,(function*(){if(!e)throw Error("No Configuration options provided, aborting");try{(0,a.CreateSessionState)();let r=(0,a.GetSessionState)();r.user.name=(0,s.ValidateUserName)(e.name),r.user.displayName=(0,s.ValidateDisplayName)(e.name),(0,a.setSessionState)(r);const n=yield(0,i.startLogin)(e.name,t),c=yield(0,o.getCredentials)(n);console.info(`Credentials created for ${e.name}`),console.log(JSON.stringify(c));const u=yield(0,i.finishLogin)({credential:c,rpOrigin:t});return r=(0,a.GetSessionState)(),r.credentials=u.result,(0,a.setSessionState)(r),u}catch(e){throw e}}))}},27:(e,t)=>{Object.defineProperty(t,"__esModule",{value:!0}),t.authenticateUserEndpoint=t.verifyAssertionEndpoint=t.assertionEndpoint=t.makeCredentialsEndpoint=t.storageKey=void 0,t.storageKey="sonr-username",t.makeCredentialsEndpoint="/name/register/start",t.assertionEndpoint="/name/register/finish",t.verifyAssertionEndpoint="/name/access/start",t.authenticateUserEndpoint="/name/access/finish"},825:function(e,t,r){var n=this&&this.__awaiter||function(e,t,r,n){return new(r||(r=Promise))((function(o,i){function s(e){try{c(n.next(e))}catch(e){i(e)}}function a(e){try{c(n.throw(e))}catch(e){i(e)}}function c(e){var t;e.done?o(e.value):(t=e.value,t instanceof r?t:new r((function(e){e(t)}))).then(s,a)}c((n=n.apply(e,t||[])).next())}))};Object.defineProperty(t,"__esModule",{value:!0}),t.getCredentials=t.createCredentials=void 0;const o=r(629),i=r(930);t.createCredentials=function(e){return n(this,void 0,void 0,(function*(){try{const t=(0,i.detectWebAuthnSupport)();if(t==o.BrowserSupport.NonHttps||t==o.BrowserSupport.Unsupported)throw new Error("Browser does not support credentials");return yield navigator.credentials.create({publicKey:e})}catch(e){throw console.error(`Error while creating public key credentials ${e.message}`),e}}))},t.getCredentials=function(e){return n(this,void 0,void 0,(function*(){if(!e)return null;console.log({pk:e});try{const t=(0,i.detectWebAuthnSupport)();if(t==o.BrowserSupport.NonHttps||t==o.BrowserSupport.Unsupported)throw new Error("Browser does not support credentials");return yield navigator.credentials.get({publicKey:e})}catch(e){throw console.error(`Error while getting public key credentials ${e.message}`),e}}))}},629:(e,t)=>{var r;Object.defineProperty(t,"__esModule",{value:!0}),t.BrowserSupport=void 0,(r=t.BrowserSupport||(t.BrowserSupport={}))[r.Supported=0]="Supported",r[r.Unsupported=1]="Unsupported",r[r.NonHttps=2]="NonHttps"},920:function(e,t,r){var n=this&&this.__createBinding||(Object.create?function(e,t,r,n){void 0===n&&(n=r);var o=Object.getOwnPropertyDescriptor(t,r);o&&!("get"in o?!t.__esModule:o.writable||o.configurable)||(o={enumerable:!0,get:function(){return t[r]}}),Object.defineProperty(e,n,o)}:function(e,t,r,n){void 0===n&&(n=r),e[n]=t[r]}),o=this&&this.__exportStar||function(e,t){for(var r in e)"default"===r||Object.prototype.hasOwnProperty.call(t,r)||n(t,e,r)};Object.defineProperty(t,"__esModule",{value:!0}),o(r(966),t),o(r(85),t)},966:function(e,t,r){var n=this&&this.__awaiter||function(e,t,r,n){return new(r||(r=Promise))((function(o,i){function s(e){try{c(n.next(e))}catch(e){i(e)}}function a(e){try{c(n.throw(e))}catch(e){i(e)}}function c(e){var t;e.done?o(e.value):(t=e.value,t instanceof r?t:new r((function(e){e(t)}))).then(s,a)}c((n=n.apply(e,t||[])).next())}))};Object.defineProperty(t,"__esModule",{value:!0}),t.startUserRegistration=void 0;const o=r(825),i=r(454),s=r(441),a=r(422);t.startUserRegistration=function(e,t=""){return n(this,void 0,void 0,(function*(){if(!e)throw Error("No Configuration options provided, aborting");try{(0,s.CreateSessionState)();let r=(0,s.GetSessionState)();r.user.name=(0,a.ValidateUserName)(e.name),r.user.displayName=(0,a.ValidateDisplayName)(e.name),(0,s.setSessionState)(r);const n=yield(0,i.startRegistration)(e.name,t),c=yield(0,o.createCredentials)(n);console.info(`Credentials created for ${e.name}`),console.log(c);const u=yield(0,i.finishRegistration)(c,t);return r=(0,s.GetSessionState)(),r.credentials=u.result,(0,s.setSessionState)(r),u}catch(e){throw console.error(`Error while registering endpoint: ${e}`),e}}))}},441:(e,t,r)=>{Object.defineProperty(t,"__esModule",{value:!0}),t.setSessionState=t.GetSessionState=t.CreateSessionState=void 0;const n=r(27);t.CreateSessionState=function(){if(!e){e={user:{name:"testuser@example.com",displayName:"testuser",id:void 0},credentials:void 0};var e=JSON.stringify(e);sessionStorage&&sessionStorage.setItem(n.storageKey,e)}},t.GetSessionState=function(){const e=(null===sessionStorage||void 0===sessionStorage?void 0:sessionStorage.getItem(n.storageKey))||"{}";return JSON.parse(e)},t.setSessionState=function(e){const t=JSON.stringify(e);sessionStorage&&sessionStorage.setItem(n.storageKey,t)}},290:(e,t)=>{var r;Object.defineProperty(t,"__esModule",{value:!0}),t.Status=void 0,(r=t.Status||(t.Status={}))[r.success=0]="success",r[r.notFound=1]="notFound",r[r.error=-1]="error"},930:(e,t,r)=>{Object.defineProperty(t,"__esModule",{value:!0}),t.buffer2string=t.bufferDecode=t.bufferEncode=t.string2buffer=t.decodeCredentialsFromAssertion=t.encodeCredentialsForAssertion=t.detectWebAuthnSupport=t.createAuthenicator=t.createAssertion=t.getStorageKey=void 0;const n=r(629),o=r(27);function i(e){try{return btoa(String.fromCharCode(...new Uint8Array(e))).replace(/\+/g,"-").replace(/\//g,"_").replace(/=/g,"")}catch(e){console.log(`Error while encoding key credentials: ${e.message}`)}}function s(e){return Uint8Array.from(atob(e),(e=>e.charCodeAt(0)))}t.getStorageKey=function(){return o.storageKey},t.createAssertion=function(e){return e?{id:e.id,rawId:i(e.rawId),type:e.type,response:{attestationObject:i(e.response.attestationObject),clientDataJSON:i(e.response.clientDataJSON)}}:{}},t.createAuthenicator=function(e){return e?{id:e.id,rawId:i(e.rawId),type:e.type,response:{authenticatorData:i(e.response.authenticatorData),clientDataJSON:i(e.response.clientDataJSON),signature:i(e.response.signature),userHandle:i(e.response.userHandle)}}:{}},t.detectWebAuthnSupport=function(){return void 0===window.PublicKeyCredential||"function"!=typeof window.PublicKeyCredential?"http:"===window.location.protocol&&"localhost"!==window.location.hostname&&"127.0.0.1"!==window.location.hostname?n.BrowserSupport.NonHttps:n.BrowserSupport.Unsupported:n.BrowserSupport.Supported},t.encodeCredentialsForAssertion=function(e){try{return{authData:new Uint8Array(e.authenticatorData),clientDataJSON:new Uint8Array(e.clientDataJSON),rawId:new Uint8Array(e.rawId),sig:new Uint8Array(e.signature),userHandle:new Uint8Array(e.response.userHandle)}}catch(e){console.error(`Error while encoding credential assertion: ${e.message}`)}},t.decodeCredentialsFromAssertion=function(e){return!!e.publicKey&&(e.publicKey.challenge=s(e.publicKey.challenge),e.publicKey.allowCredentials.forEach((function(e){e.id=s(e.id)})),!0)},t.string2buffer=function(e){return new Uint8Array(e.length).map((function(t,r){return e.charCodeAt(r)}))},t.bufferEncode=i,t.bufferDecode=s,t.buffer2string=function(e){let t="";return e.constructor!==Uint8Array&&(e=new Uint8Array(e)),e.map((function(e){return t+=String.fromCharCode(e)})),t}},454:function(e,t,r){var n=this&&this.__awaiter||function(e,t,r,n){return new(r||(r=Promise))((function(o,i){function s(e){try{c(n.next(e))}catch(e){i(e)}}function a(e){try{c(n.throw(e))}catch(e){i(e)}}function c(e){var t;e.done?o(e.value):(t=e.value,t instanceof r?t:new r((function(e){e(t)}))).then(s,a)}c((n=n.apply(e,t||[])).next())}))};Object.defineProperty(t,"__esModule",{value:!0}),t.finishLogin=t.finishRegistration=t.startLogin=t.startRegistration=t.getCredentials=t.checkUserExists=void 0;const o=r(27),i=r(441),s=r(290),a=r(930);t.checkUserExists=function(){return new Promise(((e,t)=>{try{const t=(0,i.GetSessionState)();t&&t.user.name||e(!1),fetch("/user/"+t.user.name+"/exists").then((function(t){e(!0)})).catch((function(){e(!1)}))}catch(e){console.log(`Error while validating user: ${e.message}`)}}))},t.getCredentials=function(){return new Promise(((e,t)=>{try{const r=(0,i.GetSessionState)();fetch("/credential/"+r.user.name).then((function(t){console.log(t),e(t)})).catch((function(e){console.log(`Error while resolving user credenitals for ${r.user.name}`),t()}))}catch(e){console.log(`Error while resolving user credentials for ${e.message}`)}}))},t.startRegistration=function(e,t=""){return n(this,void 0,void 0,(function*(){const r=`${t}${o.makeCredentialsEndpoint}`,n=(0,i.GetSessionState)();n.user.name=e,(0,i.setSessionState)(n);try{const e=yield fetch(r+"/"+n.user.name,{method:"GET"});if(!e||null==e)return;const t=yield null==e?void 0:e.text(),o=JSON.parse(t);if(console.log(`Credential Creation Options: ${o}`),o.publicKey&&(o.publicKey.challenge=(0,a.bufferDecode)(o.publicKey.challenge),o.publicKey.user.id=(0,a.bufferDecode)(o.publicKey.user.id)),o.publicKey.excludeCredentials)for(var s=0;s<o.publicKey.excludeCredentials.length;s++)o.publicKey.excludeCredentials[s].id=(0,a.bufferDecode)(o.publicKey.excludeCredentials[s].id);return o.publicKey}catch(e){throw console.error(`Error while making user credentials: ${e.message}`),e}}))},t.startLogin=function(e,t=""){return n(this,void 0,void 0,(function*(){const r=`${t}${o.verifyAssertionEndpoint}`;try{const t=yield fetch(r+"/"+e,{method:"GET"});if(!t||null==t)throw new Error("Could not get challenge");const n=yield null==t?void 0:t.text(),o=JSON.parse(n);return console.log(`Credential Creation Options: ${o}`),o.publicKey&&(0,a.decodeCredentialsFromAssertion)(o),o.publicKey}catch(e){throw console.error(`Error while making user credentials: ${e.message}`),e}}))},t.finishRegistration=function(e,t=""){return new Promise(((r,c)=>{try{const u=`${t}${o.assertionEndpoint}`,l=(0,i.GetSessionState)(),d=(0,a.createAssertion)(e),f=JSON.stringify(d);d&&fetch(u+"/"+l.user.name,{credentials:"same-origin",method:"POST",body:f}).then((function(e){return n(this,void 0,void 0,(function*(){const t=yield e.text();if(e.status<200||e.status>299)throw new Error(`Error while creating credential assertion: ${t}`);const n=JSON.parse(t);(0,a.decodeCredentialsFromAssertion)(n),console.log(n),r({status:s.Status.success,result:n})}))})).catch((function(e){console.log(e.name),c({error:e,status:s.Status.error})}))}catch(e){console.log(`Error while getting credential assertion: ${e.message}`),c({error:e,status:s.Status.error})}}))},t.finishLogin=function({credential:e,rpOrigin:t=""}){return new Promise(((r,c)=>{try{const c=`${t}${o.authenticateUserEndpoint}`,u=(0,i.GetSessionState)(),l=(0,a.createAuthenicator)(e),d=JSON.stringify(l);l&&fetch(c+"/"+u.user.name,{credentials:"same-origin",method:"POST",body:d}).then((function(e){return n(this,void 0,void 0,(function*(){const t=yield e.text();if(e.status<200||e.status>299)throw new Error(`Error while creating credential assertion: ${t}`);const n=JSON.parse(t);(0,a.decodeCredentialsFromAssertion)(n),console.log(n),r({result:l,status:s.Status.success})}))})).catch((function(e){console.log(e.name),r({error:e,status:s.Status.error})}))}catch(e){console.log(`Error while getting credential assertion: ${e.message}`),c({error:e,status:s.Status.error})}}))}}},t={};return function r(n){var o=t[n];if(void 0!==o)return o.exports;var i=t[n]={exports:{}};return e[n].call(i.exports,i,i.exports,r),i.exports}(920)})()));