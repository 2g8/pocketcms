<script>
    import { replace } from "svelte-spa-router";
    import ApiClient from "@/utils/ApiClient";
    import { pocketCMSSettings } from "@/stores/app";

    handler();

    function handler() {
        if (ApiClient.authStore.isValid) {
            //判断 pocketCMSSettings.flagOnboardSkip 为 true 则跳转到 /dashboard
            if ($pocketCMSSettings.flagOnboardSkip) {
                replace("/dashboard");
            } else {
                replace("/onboard");
            }
        } else {
            ApiClient.logout();
        }
    }
</script>
