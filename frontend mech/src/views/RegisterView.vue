<script setup>
  import { Authenticator} from "@aws-amplify/ui-vue";
  import "@aws-amplify/ui-vue/styles.css";
  import { Amplify, Auth } from 'aws-amplify';
Amplify.configure({
    Auth: {       
        // REQUIRED - Amazon Cognito Region
        region: 'us-east-1',

        // OPTIONAL - Amazon Cognito User Pool ID
        userPoolId: 'us-east-1_JTHfCjuRf',

        // OPTIONAL - Amazon Cognito Web Client ID (26-char alphanumeric string)
        userPoolWebClientId: '3fvj0rheotih6u3ac1ra9tdo8g',

        // OPTIONAL - Enforce user authentication prior to accessing AWS resources or not
        mandatorySignIn: false,
      }
});
// You can get the current config object
const currentConfig = Auth.configure();
console.log(currentConfig);

const formFields = {
    signUp: {
      username:{
        order:1
      },
      email: {
        order:2,
        label: 'Email:'
      },
      shopname: {
        order: 3,
        label: 'Shop name:'
      },
      address: {
        order: 5,
        label: 'Address:'
      },
      phone_number: {
        order: 4,
        label: 'Phone No.:'
      },
      longitude:{
        order: 7,
        label: 'Longitude:'
      },
      latitude:{
        order:8,
        label: 'Latitude:'
      },
      password: {
        order: 5
      },
      confirm_password: {
        order: 6
      }
    },
}
</script>


<template>
 <authenticator :form-fields="formFields">
    <template v-slot="{ user, signOut}">
      <h1 class="text-center">Welcome to Mech Moblie, {{ user.username }}!</h1>
      <p></p>
      <div class="text-center">
        <button class="btn btn-primary" @click="getandkeeptoken(user)">Proceed to Home Page</button>
        <p></p>
        <button class="btn btn-primary" @click="signOut">Sign Out</button>
      </div>
    </template>
  </authenticator>
</template>

<script>
export default {
  data(){
    return{
      name: "",
      email: "",
      address: "",
      phone_number: "",
      latitude: "",
      longitude: "",
      data: undefined
    }
  },
  methods:{
    async getandkeeptoken(user) {
      await Auth.currentSession()
        .then(data => this.data = data.getIdToken().getJwtToken())
        .catch(err => console.log(err)).then(() => sessionStorage.setItem("jwt", this.data));

        const data = new URLSearchParams({
        shopName: user.attributes.name,
        shopTel: user.attributes.phone_number,
        shopEmail: user.attributes.email,
        shopAddress: user.attributes.address,
        lat: user.attributes.latitude,
        lng: user.attributes.longitude
      })
        this.axios.post("https://a7okax4857.execute-api.us-east-1.amazonaws.com/default/shop/create-profile",data, {
        headers:{
            Authorization: this.data 
        }}).then((response)=>{
          console.log(response);
          if (response.status == 200) {
          sessionStorage.setItem("jwt", this.data);
          sessionStorage.setItem("shopID", response.data.ID);
          }
      })
      this.axios.post("https://a7okax4857.execute-api.us-east-1.amazonaws.com/default/shop/get-profile",data, {
        headers:{
            Authorization: this.data 
        }}).then((response)=>{
          console.log(response);
          if (response.status == 200) {
          sessionStorage.setItem("jwt", this.data);
          sessionStorage.setItem("shopID", response.data.ID);
          }
          })

    }
  }
};
</script>
