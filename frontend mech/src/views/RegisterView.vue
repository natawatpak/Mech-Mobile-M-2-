<template>
  <div class="pa-6 px-lg-16 px-xl-16 mx-lg-auto mx-xl-auto justify-center align-center">
    <v-card class="pa-4 rounded-xl elevation-4">
      <v-card-title class="text-h5 pa-0">Register</v-card-title>
      <v-card-subtitle class="text-wrap pa-0 text-subtitle-1">Create shop account - easily find jobs from your customer</v-card-subtitle>
      <v-form class="py-4">
        <v-text-field 
          v-model="shopName" 
          box 
          full-width 
          single-line 
          variant="outlined"
          label="Shop name" 
          color="blue-darken-3" 
          prepend-inner-icon="mdi-store-outline"
          placeholder="First name"
          />
        <v-text-field 
          v-model="shopTel" 
          box 
          full-width 
          single-line 
          variant="outlined"
          label="Telephone number" 
          color="blue-darken-3" 
          prepend-inner-icon="mdi-phone-outline"
          placeholder="Telephone number"
        />
        <v-text-field 
          v-model="shopEmail" 
          box 
          full-width 
          single-line 
          variant="outlined"
          label="Email" 
          color="blue-darken-3" 
          prepend-inner-icon="mdi-email-outline"
          placeholder="Email"
        />
        <v-text-field 
          v-model="lat" 
          box 
          full-width 
          single-line 
          variant="outlined"
          label="Shop's latitude" 
          color="blue-darken-3" 
          prepend-inner-icon="mdi-map-marker-outline"
          placeholder="Shop's latitude"
        />
        <v-text-field 
          v-model="lng" 
          box 
          full-width 
          single-line 
          variant="outlined"
          label="Shop's longtitude" 
          color="blue-darken-3" 
          prepend-inner-icon="mdi-map-marker-outline"
          placeholder="Shop's longtitude"
        />
        <v-btn @click="submitForm()" color="blue-darken-2" block size="large" >Sign up</v-btn>
      </v-form>
    </v-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      shopName: "",
      shopTel: "",
      shopEmail: "",
      shopAddress: "",
      lat: undefined,
      lng: undefined
    };
  },
  methods: {
    submitForm() {
      const data = new URLSearchParams({
        shopName: this.shopName,
        shopTel: this.shopTel,
        shopEmail: this.shopEmail,
        shopAddress: this.shopAddress,
        lat: this.lat,
        lng: this.lng,
      });
      this.axios
        .post(this.$backendApi + "shop/create-profile", data)
        .then((response) => {
          console.log(response.data.shopID);
          sessionStorage.setItem("auth", true)
          sessionStorage.setItem("shopID", response.data.shopID)
          this.$router.push("/")
        });
    },
  },
};
</script>
