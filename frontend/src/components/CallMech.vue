<template>
  <div class="pa-5">
    <v-card text-left pa-4 d-flex justify-left align-center>
      <div style="width: auto; height: 15rem" id="map"></div>
    </v-card>
    <v-card
      variant="tonal"
      @click="locatorButtonPressed()"
      class="text-left pa-4 d-flex justify-left align-center"
    >
      <v-icon icon="mdi-pin"></v-icon>
      <section>
        <v-card-title>Current Location</v-card-title>
        <v-card-text>{{
          currentLocation.lat + ", " + currentLocation.lng
        }}</v-card-text>
      </section>
    </v-card>
    <v-spacer class="my-5"></v-spacer>

    <v-card variant="tonal" class="text-left pa-4">
      <v-row class="pl-5">
        <v-card-title>Car detail</v-card-title>
        <v-spacer></v-spacer>
        <v-btn class="mx-1" @click="selectCarModal = true"
          >choose from preset</v-btn
        >
      </v-row>
      <section>
        <v-card-text>
          Type {{ car.type }}, Brand {{ car.brand }}
          <br />
          License plate: {{ car.plate }}
        </v-card-text>
      </section>
    </v-card>

    <v-dialog v-model="selectCarModal" width="800">
      <v-card>
        <v-card-title>Car preset</v-card-title>
        <v-card-text
          v-for="(car, index) in cars"
          :key="index"
          @click="selectCar(car)"
        >
          {{ car.brand }} : {{ car.plate }}
        </v-card-text>
        <v-divider></v-divider>
        <v-card-action>
          <v-btn
            block
            class="mx-1"
            prepend-icon="mdi-plus"
            @click="dialog2 = true"
            >Add new preset</v-btn
          >
        </v-card-action>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialog2" width="800">
      <v-card>
        <v-card-title>New preset</v-card-title>
        <v-form ref="form" v-model="valid">
          <v-text-field
            v-model="newCar.type"
            class="px-4"
            label="Car type"
            required
          ></v-text-field>
          <v-text-field
            v-model="newCar.brand"
            class="px-4"
            label="Brand"
            required
          ></v-text-field>
          <v-text-field
            v-model="newCar.plate"
            class="px-4"
            label="License Plate"
            required
          ></v-text-field>
        </v-form>
        <v-row justify="end" class="pa-6">
          <v-card-action>
            <v-btn
              class="mx-1"
              variant="tonal"
              color="error"
              @click="dialog2 = false"
            >
              Cancel
            </v-btn>
            <v-btn
              class="mx-1"
              color="blue"
              variant="tonal"
              @click="addNewCar(), (dialog2 = false)"
            >
              Save
            </v-btn>
          </v-card-action>
        </v-row>
      </v-card>
    </v-dialog>

    <v-spacer class="my-5"></v-spacer>
    <v-card variant="tonal" class="text-left pa-4">
      <v-card-title>Problems</v-card-title>
      <v-card-text class="d-flex">
        <v-list>
          <v-list-item class="pa-0 ma-0">
            <v-checkbox
              v-model="problem"
              value="1"
              label="Problem1"
            ></v-checkbox>
          </v-list-item>
          <v-list-item class="pa-0 ma-0">
            <v-checkbox
              v-model="problem"
              value="3"
              label="Problem3"
            ></v-checkbox>
          </v-list-item>
        </v-list>

        <v-list>
          <v-list-item class="pa-0 ma-0">
            <v-checkbox
              v-model="problem"
              value="2"
              label="Problem2"
            ></v-checkbox>
          </v-list-item>
          <v-list-item class="pa-0 ma-0">
            <v-checkbox
              v-model="problem"
              value="4"
              label="Problem4"
            ></v-checkbox>
          </v-list-item>
        </v-list>
      </v-card-text>
    </v-card>
    <v-spacer class="my-5"></v-spacer>
    <v-card variant="tonal" class="text-left pa-4">
      <v-card-title>Details</v-card-title>
      {{description}}
      <v-textarea
        outlined
        name="description"
        v-model="description"
        rows="5"
        label="Description"
        bg-color="white"
      ></v-textarea>
    </v-card>
    <v-spacer class="my-5"></v-spacer>

    <v-spacer class="my-5"></v-spacer>
    <section class="text-center">
      <div class="text-decoration-none" @click= "addTicket()"
        ><v-btn class="mx-1" variant="tonal" color="blue" 
          >Find Service</v-btn
        ></div
      >
    </section>
  </div>
</template>

<script>
/* eslint-disable */
export default {
  data() {
    return {
      map: null,
      currentLocation: { lat: 0, lng: 0 },
      car: { id: "", type: "", brand: "", plate: "" },
      cars: [],
      selectCarModal: false,
      dialog2: false,
      newCar: {
        type: "",
        brand: "",
        plate: "",
      },
      problem: [],
      description: "",
      files: [],
    };
  },
  mounted() {
    this.initMap();
    this.locatorButtonPressed();
    this.setMarker(this.mapCenter, "A");
    this.getCarList();
  },
  methods: {
    locatorButtonPressed() {
      const success = (position) => {
        const latitude = position.coords.latitude;
        const longitude = position.coords.longitude;
        // Do something with the position
        console.log(latitude + ", " + longitude);
        this.currentLocation.lat = latitude;
        this.currentLocation.lng = longitude;
        this.setMarker(this.currentLocation);
        this.map.setCenter(this.currentLocation);
      };

      const error = (err) => {
        console.log(err);
      };

      // This will open permission popup
      navigator.geolocation.getCurrentPosition(success, error);
    },
    initMap() {
      this.map = new google.maps.Map(document.getElementById("map"), {
        center: this.currentLocation,
        zoom: 15,
        maxZoom: 20,
        minZoom: 3,
        streetViewControl: false,
        mapTypeControl: false,
        fullscreenControl: false,
        zoomControl: true,
      });
    },
    setMarker(Points) {
      const markers = new google.maps.Marker({
        position: Points,
        map: this.map,
      });
    },
    selectCar(car) {
      this.car = car;
      this.selectCarModal = false;
    },
    getCarList() {
      const data = new URLSearchParams({
        cusID: sessionStorage.getItem("cusID"),
      });
      console.log(sessionStorage.getItem("cusID"));

      this.axios
        .post(this.$backendApi + "customer/get-car-list", data)
        .then((response) => {
          console.log(response.data);
          this.cars = response.data;
        });
    },
    addTicket() {
      const data = new URLSearchParams({
        cusID: sessionStorage.getItem("cusID"),
        carID: this.car.id,
        problem: this.problem,
        status: "Active",
        lng: this.currentLocation.lng,
        lat: this.currentLocation.lat,
        description: this.description,
      });
      this.axios
        .post(this.$backendApi + "customer/add-ticket", data)
        .then((response) => {
          console.log(response.data);
          sessionStorage.setItem("ticketID", response.data.ticketID);
        }).then(()=>{
          this.$router.push('/progress');
        });
    },
    addNewCar() {
      console.log("clicked")
      const data = new URLSearchParams({
        cusID: sessionStorage.getItem("cusID"),
        plateNum: this.newCar.plate,
        issuedAt: this.newCar.issuedAt,
        color: this.newCar.color,
        type: this.newCar.type,
        brand: this.newCar.brand,
        build: this.newCar.build,
      });
      this.axios
        .post(this.$backendApi + "customer/add-car", data)
        .then((response) => {
          console.log(response.data);
          this.getCarList();
        });
    },
  },
};
</script>