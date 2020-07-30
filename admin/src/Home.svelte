<script>
  let loadingFeed = true;
  let feed = undefined;
  let errorMessage = undefined;
  import { onMount } from "svelte";
  onMount(get_feed);
  async function get_feed() {
    let url = "https://raw.githubusercontent.com/pagopa/io-sdk/master/feed.json";
      fetch(url)
      .then(async res => {
        if (res.ok) {
          feed = await res.json();
          loadingFeed = false;
        } else {
          errorMessage = { error: "404 Not Found" };
          loadingFeed = false;
        }
      })
      .catch(err => {
        errorMessage = { error: err };
        console.log(errorMessage);
        loadingFeed = false;
      });
  }
</script>

<div class="it-hero-wrapper">
  <div class="container-fluid">
    <div class="row">
        <div class="col-9">
          <div class="it-hero-text-wrapper bg-dark">
              <span class="it-category">Welcome</span>
              <h1 class="no_toc">IO-SDK</h1>
              <p class="d-none d-lg-block">Welcome to the development kit to easily build IO importers.</p>
          </div>
        </div>
    </div>
    <div class="row">
        <div class="col-12">
          <div class="it-hero-text-wrapper bg-dark">
            <span class="it-category">Ultime novità</span>
            {#if loadingFeed}
              <div>loading...</div>
            {:else if errorMessage}
                <div>
                  <span>Si è verificato un errore: {errorMessage.error}</span>
                </div>
            {:else if feed && feed.items}
              {#each feed.items as item}
                <div style="width:90%;color:white">
                  <h3>{item.title}</h3>
                  {#if item.summary}
                    {item.summary}
                  {:else if item.content_html}
                    {@html item.content_html}
                  {:else}
                    {item.content_text}
                  {/if}
                  {#if item.url}
                    <div style="width:100%;text-align:right">
                      <a style="color:white" href="{item.url}">leggi la news</a>
                    </div>
                  {/if}
                </div>
              {/each}
            {:else}
                <div>
                  <span>Nessuna novità</span>
                </div>
            {/if}
          </div>
        </div>
    </div>
  </div>
</div>