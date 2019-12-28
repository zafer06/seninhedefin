<template>
  <section class="targetList text-center">
    <div class="container">
      <div class="title">
        <h1> Senin Hedefin </h1>
      </div>

      <div class="row mt-5">
        <div class="col-lg-6 col-md-6 col-sm-6 col-12 mx-auto">
          <form>
            <div class="row">
              <div class="col-lg-6 col-md-6 col-sm-6 col-12">
                <input
                type="text"
                class="form-control nameAndMail"
                placeholder="Ad Soyad"
                required
                >
              </div>
              <div class="col-lg-6 col-md-6 col-sm-6 col-12">
                <input
                type="text"
                class="form-control nameAndMail"
                placeholder="E-mail"
                required
                >
              </div>
              <div class="col-lg-12 col-md-12 col-sm-12 col-12">
                <input
                    @keyup.enter="addTarget"
                    type="text"
                    placeholder="Senin Hedefin Nedir?"
                    autofocus="autofocus"
                    autocomplate="off"
                    class="newTarget form-control"
                >
                <transition name="error">
                  <div class="error" v-show="error">
                    <p> En fazla 10 değer girebilirsiniz. </p>
                  </div>
                </transition>
            </div>

          </div>
            <div clas="col-lg-12 col-md-12 col-sm-12 col-12">
              <transition-group name="fade" tag="div" appear>
                <div class="targetItem" v-for="target in targetItems" :key="target">
                  <span class="item"> {{target}} </span>
                  <span class="closeBtn" @click="removeTarget">x</span>
                </div>
              </transition-group>
            </div>

          </form>
        </div>
      </div>
    </div>
  </section>
</template>


<script>
export default {
  data(){
    return {
      error:false,
      targetItems:["hello","word"]
    }
  },
  methods:{
    addTarget(){
      const valueText = event.target;
      if (this.targetItems.length <= 4) {
        this.targetItems.push(valueText.value);
        valueText.value='';
      } else {
        this.error = true;
        setTimeout(() => {
          this.error = false;
        },1500);

      }
    },
    removeTarget(index){
      this.targetItems.splice(index, 1);
    }
  }
}

</script>
