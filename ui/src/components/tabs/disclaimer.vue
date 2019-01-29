<template>
	<div>
		<h4>Disclaimer</h4>
		<p class="code">
			We are not responsible for any loss: <b>Ethereum</b>, {{app}} and some of the underlying script libraries (web3js, etc) and software (go-ethereum, etc) we use are under active development. We have tested {{app}} successfully, however there is always the chance of something going wrong and unexpected events can occur, and might lead to ETH lost. We do not recommend the usage of {{app}} for crypto operations (hardware based security is always better, aka, cold wallets). But if you want to use it, feel free to do it so, but rebember the facts that online website can be compromised, accounts stolen by fake sites, system hacked by cybercriminals, etc. In short, please do not transfer more than what you are willing to lose, and please be careful. If something happens, we are sorry, but we are not responsible for the lost Ether.
		</p>
		<div class="copyright">
			Copyright © 2019 {{app}}
		</div>
		<p class="code">
			THE SERVICE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SERVICE.
		</p>
	</div>
</template>

<script>
export default {
  name: 'disclaimer-tab',
  data () {
    return {
    	app: process.env.APP_NAME
    }
  },
  methods: {
  },
  created(){
    log("disclaimer-tab::created");
  },
  mounted(){
    log("disclaimer-tab::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.copyright {
	font-family: 'IBM Plex Mono';
	font-weight: bold;
	padding-bottom: 10px;
}
.code {
	font-family: 'IBM Plex Mono';
}
h4, h5, p {
	text-align: justify;
}
</style>