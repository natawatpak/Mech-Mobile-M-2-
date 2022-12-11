<script setup>
  import { Authenticator } from "@aws-amplify/ui-vue";
  import "@aws-amplify/ui-vue/styles.css";
  import { Amplify, Auth } from 'aws-amplify';
Amplify.configure({
    Auth: {       
        // REQUIRED - Amazon Cognito Region
        region: 'us-east-1',

        // OPTIONAL - Amazon Cognito User Pool ID
        userPoolId: 'us-east-1_qWVv2xtlo',

        // OPTIONAL - Amazon Cognito Web Client ID (26-char alphanumeric string)
        userPoolWebClientId: '6oa2an34ok543djt1h5pg7c3v',

        // OPTIONAL - Enforce user authentication prior to accessing AWS resources or not
        mandatorySignIn: false,
      }
});
// You can get the current config object
const currentConfig = Auth.configure();
console.log(currentConfig);
</script>


<template>
 <authenticator :sign-up-attributes="[
    'email',
    'family_name',
    'given_name',
    'phone_number' ]">
    <template v-slot="{ user, signOut }">
      <h1>Hello {{ user.username }}!</h1>
      <button @click="signOut">Sign Out</button>
      <button @click="getandkeeptoken">Get JWT Token</button>
      <button @click="submitForm(user)">test API call with Cognito</button>
      <p>{{data}}</p>
      <p>{{user.attributes.email}}</p>
      <p>{{user.attributes.given_name}}</p>
      <p>{{user.attributes.family_name}}</p>
      <p>{{user.attributes.phone_number}}</p>
    </template>
  </authenticator>
</template>

<script>
export default {
  data(){
    return{
      fName: "",
      lName: "",
      tel: "",
      email: "",
      data: ""
    }
  },
  methods:{
    submitForm(user){
      const data = new URLSearchParams({
        fName: user.attributes.given_name,
        lName: user.attributes.family_name,
        tel: user.attributes.phone_number,
        email: user.attributes.email,
      })
      this.axios.post("https://a7okax4857.execute-api.us-east-1.amazonaws.com/default/customer/create-profile",data, {
        headers:{
            Authorization: this.data 
        }}).then((response)=>{
          console.log(response);
          if (response.status == 200) {
            console.log("GGG");
          sessionStorage.setItem("jwt", this.data);
          sessionStorage.setItem("cusID", response.data.ID);
          sessionStorage.setItem("fName", response.data.fName);
          sessionStorage.setItem("lName", response.data.lName);
          sessionStorage.setItem("tel", response.data.tel);
          sessionStorage.setItem("email", response.data.email);
          }
      })
    },
    getandkeeptoken() {
      Auth.currentSession()
        .then(data => this.data = data.getIdToken().getJwtToken())
        .catch(err => console.log(err)).then(() => sessionStorage.setItem("jwt", this.data));
    }
  }
};
</script>
