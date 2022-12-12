<template>
  <div class="pa-4 px-lg-16 px-xl-16 mx-lg-auto mx-xl-auto text-center justify-center" v-if="ticket">
    <v-row class="ma-2 justify-space-between align-center pa-0">
      <div class="text-left text-h5 my-2 font-weight-bold">{{ticket.cus.fName}} | {{ticket.car.plate}}</div>
      <v-chip color="yellow-darken-3" text-color="white">{{ticket.status}}</v-chip>
    </v-row>

    <v-card class="text-left pa-4 rounded-lg elevation-4">
      <v-card-title class="text-h6 pa-0">Status</v-card-title>
      <v-timeline direction="horizontal" line-inset="12" class="pa-0">
      <v-timeline-item size="large" v-model="items" :dot-color="ticket.status == stage[0] || ticket.status == stage[1] || ticket.status == stage[2] || ticket.status == stage[3] || ticket.status == stage[4] ? 'green-lighten-1':'white'" :icon=items[0]>
        <template v-slot:opposite>
          Accept
        </template>
       </v-timeline-item>

      <v-timeline-item size="large" v-model="items" :dot-color="ticket.status == stage[1] || ticket.status == stage[2] || ticket.status == stage[3] || ticket.status == stage[4]? 'green-lighten-1':'white'" :icon=items[1]>
      <template v-slot:opposite>
      </template>
        On the way
      </v-timeline-item>

      <v-timeline-item size="large" v-model="items" :dot-color="ticket.status == stage[2] || ticket.status == stage[3] || ticket.status == stage[4]? 'green-lighten-1':'white'" :icon=items[2]>
        <template v-slot:opposite>
          On process
        </template>
      </v-timeline-item>

      <v-timeline-item size="large" v-model="items" :dot-color="ticket.status == stage[3] || ticket.status == stage[4]? 'green-lighten-1':'white'" :icon=items[3]>
        <template v-slot:opposite>
        </template>
        Finish
      </v-timeline-item>
      </v-timeline>
      <div>
      <v-card-subtitle v-if="(ticket.status == stage[0])" class="pa-0 py-2 text-wrap">{{otw}}</v-card-subtitle>
      <v-card-subtitle v-else-if="(ticket.status == stage[1])" class="pa-0 py-2 text-wrap">{{onp}}</v-card-subtitle>
      <v-card-subtitle v-else-if="(ticket.status == stage[2])" class="pa-0 py-2 text-wrap">{{fin}}</v-card-subtitle>
      </div>
      <v-row justify="end" class="pa-4">
        <v-btn :disabled="valid" class="mx-1 " variant="tonal" color="blue-darken-2" @click="checkStage">Next stage</v-btn>
      </v-row>
    </v-card>
    
    <v-spacer class="my-3"></v-spacer>

    <v-card class="text-left pa-4 d-flex justify-left align-center rounded-lg elevation-4">
      <section>
        <v-card-title class="text-h6 pa-0">Customer name</v-card-title>
        <v-card-text class="pa-0 py-2 text-subtitle-1">
          {{(ticket.cus.fName+' '+ticket.cus.lName)}}
        </v-card-text>
      </section>
    </v-card>

    <v-spacer class="my-3"></v-spacer>

    <v-card class="text-left pa-4 rounded-lg elevation-4">
      <v-title class="text-h6 pa-0">Car Details</v-title>
        <v-card-text class="text-h7 pa-0 py-2 text-subtitle-1">
          Type: {{ticket.car.type}}
          <br>
          Brand: {{ticket.car.brand}}
          <br />
          License plate: {{ticket.car.plate}}
        </v-card-text>
    </v-card>

    <v-spacer class="my-3"></v-spacer>

    <v-card class="text-left pa-4 d-flex justify-left align-center rounded-lg elevation-4">
      <section>
        <v-card-title class="text-h6 pa-0">Current Location</v-card-title>
        <v-card-text class="pa-0 py-2 text-subtitle-1">
          Latitude: {{ticket.location.lat}} 
          <br>
          Longtitude: {{ticket.location.lng}}
          <br>
          Distance: {{ticket.location.distance}} km
        </v-card-text>
      </section>
    </v-card>

    <v-spacer class="my-3"></v-spacer>

    <v-card class="text-left mb-4 pa-4 rounded-lg elevation-4">
      <v-title class="text-h6 pa-0">Problems</v-title>
      <v-card-text class="text-h7 pa-0 text-subtitle-1">
        <li v-for="p in splitProblem(ticket.problem)" :key="p">{{ p }}</li>
      </v-card-text>
    </v-card>

    <v-spacer class="my-3"></v-spacer>

    <v-card v-if="checkDesc()" v-show="show" class="text-left mb-4 pa-4 rounded-lg elevation-4">
      <v-title class="text-h6 pa-0">Description</v-title>
      <v-card-text class="text-h7 pa-0">
        <p>{{ticket.description}}</p>
      </v-card-text>
    </v-card> 

    <v-dialog v-model="dialog" width="800" class="text-left pa-4">
      <v-card class="text-left pa-4">
        <v-card-title class="text-h6 pa-0">Status: {{ticket.status}}</v-card-title>   
          <v-timeline direction="horizontal" line-inset="12" class="pa-0">
            <v-timeline-item size="large" v-model="items" :dot-color="ticket.status == stage[0] || ticket.status == stage[1] || ticket.status == stage[2] || ticket.status == stage[3] || ticket.status == stage[4] ? 'green-lighten-1':'white'" :icon=items[0]>
            <template v-slot:opposite>
              Accept
            </template>
          </v-timeline-item>

          <v-timeline-item size="large" v-model="items" :dot-color="ticket.status == stage[1] || ticket.status == stage[2] || ticket.status == stage[3] || ticket.status == stage[4]? 'green-lighten-1':'white'" :icon=items[1]>
          <template v-slot:opposite>
          </template>
            On the way
          </v-timeline-item>

          <v-timeline-item size="large" v-model="items" :dot-color="ticket.status == stage[2] || ticket.status == stage[3] || ticket.status == stage[4]? 'green-lighten-1':'white'" :icon=items[2]>
            <template v-slot:opposite>
              On process
            </template>
          </v-timeline-item>

          <v-timeline-item size="large" v-model="items" :dot-color="ticket.status == stage[3] || ticket.status == stage[4]? 'green-lighten-1':'white'" :icon=items[3]>
            <template v-slot:opposite>
            </template>
            Finish
          </v-timeline-item>
        </v-timeline>
        
        <div class="pt-2">
          <v-text class="text-subtitle-1 pa-0">Please type 'confirm' to continue</v-text>
          <v-text-field v-model="confirm" label="Confirm" required></v-text-field>
        </div>

        <v-row justify="end" class="pa-4">
            <v-card-action>
              <v-btn class="mx-1" variant="tonal" color="error" @click="dialog = false">
                Cancel
              </v-btn>

              <v-btn :disabled="valid2" class="mx-1" color="blue-darken-2" variant="tonal" @click="checkConfirm()">
                Confirm
              </v-btn>
            </v-card-action>
          </v-row>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialog2" width="800" class="text-left pa-4">
      <v-card class="text-left pa-4">
        <v-card-title class="text-h5">Status {{ticket.status}}</v-card-title> 
          <v-timeline direction="horizontal" line-inset="12" class="pa-0">
            <v-timeline-item size="large" v-model="items" :dot-color="ticket.status == stage[0] || ticket.status == stage[1] || ticket.status == stage[2] || ticket.status == stage[3] || ticket.status == stage[4] ? 'green-lighten-1':'white'" :icon=items[0]>
              <template v-slot:opposite>
                Accept
              </template>
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="ticket.status == stage[1] || ticket.status == stage[2] || ticket.status == stage[3] || ticket.status == stage[4]? 'green-lighten-1':'white'" :icon=items[1]>
            <template v-slot:opposite>
            </template>
              On the way
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="ticket.status == stage[2] || ticket.status == stage[3] || ticket.status == stage[4]? 'green-lighten-1':'white'" :icon=items[2]>
              <template v-slot:opposite>
                On process
              </template>
            </v-timeline-item>

            <v-timeline-item size="large" v-model="items" :dot-color="ticket.status == stage[3] || ticket.status == stage[4]? 'green-lighten-1':'white'" :icon=items[3]>
              <template v-slot:opposite>
              </template>
              Finish
            </v-timeline-item>
          </v-timeline>

          <div class="pt-2 pa-0 ma-0">
            <v-radio-group column v-model="finishOption">
              <v-radio label="All complete" value="Finish:Completed"></v-radio>
              <v-radio label="Go to garage" value="Finish:Garage"></v-radio>
            </v-radio-group>
          </div>

          <div class="pa-0">
            <v-text class="text-subtitle-1 pa-0">Please type 'confirm' to continue</v-text>
            <v-text-field v-model="confirm" label="Confirm" required></v-text-field>
          </div>

        <v-row justify="end" class="pa-4">
          <v-card-action>
            <v-btn class="mx-1" variant="tonal" color="error" @click="dialog2 = false">
              Cancel
            </v-btn>

            <v-btn class="mx-1" color="blue" variant="tonal" @click="checkConfirm()">
              Submit
            </v-btn>
          </v-card-action>
        </v-row>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      show: true,
      ticket: null,
      finishOption: null,
      dialog: false,
      dialog2: false,
      stage: [ 'Accepted', 'On the way', 'Processing', 'Finish:Garage', 'Finish:Completed'],
      items: ['mdi-thumb-up', 'mdi-car', 'mdi-wrench', 'mdi-check'],
      otw: "Next stage: 'On the way'. Are your mechanic ready to go?",
      onp: "Next stage: 'On process'. Your mechanic is fixing customer car.",
      fin: "Last stage: 'Finish'. All completed!"
    };
  },
  mounted() {
    this.getTicket(sessionStorage.getItem("ticketID"))

    this.socket = new WebSocket(this.$wsApi + "shop/ws/"+sessionStorage.getItem("ticketID"));
    console.log("Attempting Connection...");

    this.socket.onopen = () => {
      console.log("Successfully Connected");
      this.socket.send("Accepted");
    };

    this.socket.addEventListener("message", (event) => {
      console.log("Message from server ", event.data);
    });

    this.socket.onclose = (event) => {
      console.log("Socket Closed Connection: ", event);
      this.socket.send("Client Closed!");
    };

    this.socket.onerror = (error) => {
      this.console.log("Socket Error: ", error);
    };

  },
  methods: {
    splitProblem(p) {
      return (this.problems = p.split(","));
    },
    checkDesc() {
      if (this.ticket.description == '') {
        this.show = false;
      }
    },
    goHome() {
      this.$route.router.go('/');
    },
    checkConfirm() {
      if (this.confirm == 'confirm') {
        if (this.ticket.status == this.stage[0]) {
          this.dialog = false;
          this.confirm = '';
          this.updateTicketStatus('On the way')
        }
        else if (this.ticket.status == this.stage[1]) {
          this.dialog = false;
          this.confirm = '';
          this.updateTicketStatus('Processing')
        }
        else if (this.ticket.status == this.stage[2]) {
          this.dialog2 = false;
          this.confirm = '';
          this.updateTicketStatus(this.finishOption)
          this.goHome();
        }
      }
    },
    checkStage() {
      if (this.ticket.status == this.stage[0] || this.ticket.status == this.stage[1]) {
        this.dialog = true;      
      }
      else if (this.ticket.status == this.stage[2]) {
        this.dialog2 = true;
      }
      else {
        this.valid = true;
      }
    },
    getTicket(id) {
      const data = new URLSearchParams({
        ticketID: id,
      });

      this.axios
        .post(this.$backendApi + "shop/get-ticket",data)
        .then((response) => {
          console.log(response.data);
          this.ticket = response.data;
          Object.keys(response.data).forEach((prop)=> console.log(prop));
        });
    },
    updateTicketStatus(status){
      this.socket.send(status);

      const data = new URLSearchParams({
        ticketID: this.ticket.ticketID,
        cusID: this.ticket.cus.cusID,
        carID: this.ticket.car.carID,
        problem: this.ticket.problem,
        shopID: sessionStorage.getItem("shopID"),
        status: status,
        lat: this.ticket.location.lat,
        lng: this.ticket.location.lng,
        description: this.ticket.description
      });

      this.axios
        .post(this.$backendApi + "shop/update-ticket", data)
        .then((response) => {
          console.log(response);
          this.ticket.status = response.data.status
        })
    },
  },
};
</script>
