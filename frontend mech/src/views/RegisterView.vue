<template>
  <div>
    Register
    <v-form>
      <v-text-field
        v-model="shopName"
        placeholder="please input shop name"
      />
      <v-text-field
        v-model="shopTel"
        placeholder="please input your shop telephone number"
      />
      <v-text-field v-model="shopEmail" placeholder="shopEmail" />
      <v-text-field
        v-model="lat"
        placeholder="please input your lat"
      />
      <v-text-field
        v-model="lng"
        placeholder="please input your lng"
      />

      <v-btn @click="submitForm()">submit</v-btn>
    </v-form>
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
