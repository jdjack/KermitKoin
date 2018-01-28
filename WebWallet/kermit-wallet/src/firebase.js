import firebase from 'firebase'

var config = {
  apiKey: "AIzaSyDk-Rga8m0rqrxOqAKUjfXTw1eyhf9DdhE",
  authDomain: "kermit-koin.firebaseapp.com",
  databaseURL: "https://kermit-koin.firebaseio.com",
  projectId: "kermit-koin",
  storageBucket: "",
  messagingSenderId: "762618241145"
};
firebase.initializeApp(config);


export default firebase;
