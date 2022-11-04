// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Order;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.ovh.Order.inputs.GetCartArgs;
import com.pulumi.ovh.Order.inputs.GetCartPlainArgs;
import com.pulumi.ovh.Order.inputs.GetCartProductArgs;
import com.pulumi.ovh.Order.inputs.GetCartProductOptionsArgs;
import com.pulumi.ovh.Order.inputs.GetCartProductOptionsPlainArgs;
import com.pulumi.ovh.Order.inputs.GetCartProductOptionsPlanArgs;
import com.pulumi.ovh.Order.inputs.GetCartProductOptionsPlanPlainArgs;
import com.pulumi.ovh.Order.inputs.GetCartProductPlainArgs;
import com.pulumi.ovh.Order.inputs.GetCartProductPlanArgs;
import com.pulumi.ovh.Order.inputs.GetCartProductPlanPlainArgs;
import com.pulumi.ovh.Order.outputs.GetCartProductInvokeResult;
import com.pulumi.ovh.Order.outputs.GetCartProductOptionsInvokeResult;
import com.pulumi.ovh.Order.outputs.GetCartProductOptionsPlanResult;
import com.pulumi.ovh.Order.outputs.GetCartProductPlanResult;
import com.pulumi.ovh.Order.outputs.GetCartResult;
import com.pulumi.ovh.Utilities;
import java.util.concurrent.CompletableFuture;

public final class OrderFunctions {
    /**
     * Use this data source to create a temporary order cart to retrieve information order cart products.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .description(&#34;...&#34;)
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetCartResult> getCart(GetCartArgs args) {
        return getCart(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to create a temporary order cart to retrieve information order cart products.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .description(&#34;...&#34;)
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetCartResult> getCartPlain(GetCartPlainArgs args) {
        return getCartPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to create a temporary order cart to retrieve information order cart products.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .description(&#34;...&#34;)
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetCartResult> getCart(GetCartArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Order/getCart:getCart", TypeShape.of(GetCartResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to create a temporary order cart to retrieve information order cart products.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .description(&#34;...&#34;)
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetCartResult> getCartPlain(GetCartPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Order/getCart:getCart", TypeShape.of(GetCartResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information of order cart product products.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var plans = OrderFunctions.getCartProduct(GetCartProductArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .product(&#34;...&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetCartProductInvokeResult> getCartProduct(GetCartProductArgs args) {
        return getCartProduct(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information of order cart product products.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var plans = OrderFunctions.getCartProduct(GetCartProductArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .product(&#34;...&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetCartProductInvokeResult> getCartProductPlain(GetCartProductPlainArgs args) {
        return getCartProductPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information of order cart product products.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var plans = OrderFunctions.getCartProduct(GetCartProductArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .product(&#34;...&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetCartProductInvokeResult> getCartProduct(GetCartProductArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Order/getCartProduct:getCartProduct", TypeShape.of(GetCartProductInvokeResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information of order cart product products.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var plans = OrderFunctions.getCartProduct(GetCartProductArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .product(&#34;...&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetCartProductInvokeResult> getCartProductPlain(GetCartProductPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Order/getCartProduct:getCartProduct", TypeShape.of(GetCartProductInvokeResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information of order cart product options.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductOptionsArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var options = OrderFunctions.getCartProductOptions(GetCartProductOptionsArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .product(&#34;cloud&#34;)
     *             .planCode(&#34;project&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetCartProductOptionsInvokeResult> getCartProductOptions(GetCartProductOptionsArgs args) {
        return getCartProductOptions(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information of order cart product options.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductOptionsArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var options = OrderFunctions.getCartProductOptions(GetCartProductOptionsArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .product(&#34;cloud&#34;)
     *             .planCode(&#34;project&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetCartProductOptionsInvokeResult> getCartProductOptionsPlain(GetCartProductOptionsPlainArgs args) {
        return getCartProductOptionsPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information of order cart product options.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductOptionsArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var options = OrderFunctions.getCartProductOptions(GetCartProductOptionsArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .product(&#34;cloud&#34;)
     *             .planCode(&#34;project&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetCartProductOptionsInvokeResult> getCartProductOptions(GetCartProductOptionsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Order/getCartProductOptions:getCartProductOptions", TypeShape.of(GetCartProductOptionsInvokeResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information of order cart product options.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductOptionsArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var options = OrderFunctions.getCartProductOptions(GetCartProductOptionsArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .product(&#34;cloud&#34;)
     *             .planCode(&#34;project&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetCartProductOptionsInvokeResult> getCartProductOptionsPlain(GetCartProductOptionsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Order/getCartProductOptions:getCartProductOptions", TypeShape.of(GetCartProductOptionsInvokeResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information of order cart product options plan.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductOptionsPlanArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var plan = OrderFunctions.getCartProductOptionsPlan(GetCartProductOptionsPlanArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .priceCapacity(&#34;renew&#34;)
     *             .product(&#34;cloud&#34;)
     *             .planCode(&#34;project&#34;)
     *             .optionsPlanCode(&#34;vrack&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetCartProductOptionsPlanResult> getCartProductOptionsPlan(GetCartProductOptionsPlanArgs args) {
        return getCartProductOptionsPlan(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information of order cart product options plan.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductOptionsPlanArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var plan = OrderFunctions.getCartProductOptionsPlan(GetCartProductOptionsPlanArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .priceCapacity(&#34;renew&#34;)
     *             .product(&#34;cloud&#34;)
     *             .planCode(&#34;project&#34;)
     *             .optionsPlanCode(&#34;vrack&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetCartProductOptionsPlanResult> getCartProductOptionsPlanPlain(GetCartProductOptionsPlanPlainArgs args) {
        return getCartProductOptionsPlanPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information of order cart product options plan.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductOptionsPlanArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var plan = OrderFunctions.getCartProductOptionsPlan(GetCartProductOptionsPlanArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .priceCapacity(&#34;renew&#34;)
     *             .product(&#34;cloud&#34;)
     *             .planCode(&#34;project&#34;)
     *             .optionsPlanCode(&#34;vrack&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetCartProductOptionsPlanResult> getCartProductOptionsPlan(GetCartProductOptionsPlanArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Order/getCartProductOptionsPlan:getCartProductOptionsPlan", TypeShape.of(GetCartProductOptionsPlanResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information of order cart product options plan.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductOptionsPlanArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var plan = OrderFunctions.getCartProductOptionsPlan(GetCartProductOptionsPlanArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .priceCapacity(&#34;renew&#34;)
     *             .product(&#34;cloud&#34;)
     *             .planCode(&#34;project&#34;)
     *             .optionsPlanCode(&#34;vrack&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetCartProductOptionsPlanResult> getCartProductOptionsPlanPlain(GetCartProductOptionsPlanPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Order/getCartProductOptionsPlan:getCartProductOptionsPlan", TypeShape.of(GetCartProductOptionsPlanResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information of order cart product plan.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductPlanArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var plan = OrderFunctions.getCartProductPlan(GetCartProductPlanArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .priceCapacity(&#34;renew&#34;)
     *             .product(&#34;cloud&#34;)
     *             .planCode(&#34;project&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetCartProductPlanResult> getCartProductPlan(GetCartProductPlanArgs args) {
        return getCartProductPlan(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information of order cart product plan.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductPlanArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var plan = OrderFunctions.getCartProductPlan(GetCartProductPlanArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .priceCapacity(&#34;renew&#34;)
     *             .product(&#34;cloud&#34;)
     *             .planCode(&#34;project&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetCartProductPlanResult> getCartProductPlanPlain(GetCartProductPlanPlainArgs args) {
        return getCartProductPlanPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to retrieve information of order cart product plan.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductPlanArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var plan = OrderFunctions.getCartProductPlan(GetCartProductPlanArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .priceCapacity(&#34;renew&#34;)
     *             .product(&#34;cloud&#34;)
     *             .planCode(&#34;project&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetCartProductPlanResult> getCartProductPlan(GetCartProductPlanArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ovh:Order/getCartProductPlan:getCartProductPlan", TypeShape.of(GetCartProductPlanResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to retrieve information of order cart product plan.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ovh.Order.OrderFunctions;
     * import com.pulumi.ovh.Order.inputs.GetCartArgs;
     * import com.pulumi.ovh.Order.inputs.GetCartProductPlanArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
     *             .ovhSubsidiary(&#34;fr&#34;)
     *             .description(&#34;my cart&#34;)
     *             .build());
     * 
     *         final var plan = OrderFunctions.getCartProductPlan(GetCartProductPlanArgs.builder()
     *             .cartId(mycart.applyValue(getCartResult -&gt; getCartResult.id()))
     *             .priceCapacity(&#34;renew&#34;)
     *             .product(&#34;cloud&#34;)
     *             .planCode(&#34;project&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetCartProductPlanResult> getCartProductPlanPlain(GetCartProductPlanPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ovh:Order/getCartProductPlan:getCartProductPlan", TypeShape.of(GetCartProductPlanResult.class), args, Utilities.withVersion(options));
    }
}
