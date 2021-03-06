
import java.io.IOException;
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@FFF
public class Main extends AbstractMain<String> implements InterfaceMain {

    @FFF
    public static final Integer INT_MAX = Integer.MAX_VALUE;
    public static final Integer INT_MIN = Integer.MIN_VALUE;

    public static final Long LONG_MAX = Long.MAX_VALUE;
    public static final Long LONG_MIN = Long.MIN_VALUE;

    public static final Double DOUBLE_MAX = Double.MAX_VALUE;
    public static final Double DOUBLE_MIN = Double.MIN_VALUE;

    public static final Float FLOAT_MAX = Float.MAX_VALUE;
    public static final Float FLOAT_MIN = Float.MIN_VALUE;

    public static final Short SHORT_MAX = Short.MAX_VALUE;
    public static final Short SHORT_MIN = Short.MIN_VALUE;

    private String name;

    public Main() {
    }

    public Main(String name) {
        this.name = name;
    }

    public static void main(String[] args) throws IOException {

        System.out.println("Hello word!");
        Main main = new Main();

        String p = "777";
        Runnable r = () -> {
            System.out.println("run");
            main.hi(p);
        };

        r.run();

        int a = 1;
        String A = "";
        switch (a) {
            case 1: {
                A = "a";
            }
            case 2: {
                A = "c";
            }
            case 3: {
                A = "D";
            }
            default: {
                A = "Z";
            }
        }

        try {
            System.out.println("try");
        } catch (Exception e) {
            System.out.println("catch");
        } finally {
            System.out.println("finally");
        }
    }

    @Override
    @FFF("3FFF")
    @Name(name = "My Name", index = 666)
    @Handler(name = MainEnum.ONE, clazz = InterfaceMain.class, value = {"1", "2"})
    public void say() {
        System.out.println("ha ha ha ha");
    }

    @Override
    public String hi(String p) {
        return "Hi,2021" + p;
    }
}


interface InterfaceMain {
    void say();
}

abstract class AbstractMain<T> {
    abstract T hi(T p);
}

@Target({ElementType.METHOD, ElementType.FIELD, ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
@interface FFF {
    String value() default "FFF_DEFAULT";
}

@Target({ElementType.METHOD, ElementType.FIELD, ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
@interface Handler {

    MainEnum name() default MainEnum.DEFAULT;

    Class<InterfaceMain> clazz() default InterfaceMain.class;

    String[] value();
}

@Target({ElementType.METHOD, ElementType.FIELD, ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
@interface Name {
    String name() default "Name";

    int index() default -1;
}

enum MainEnum {

    DEFAULT("-1"),
    ONE("1"),
    TWO("2"),
    THREE("3"),

    ;
    private String name;

    MainEnum(String name) {
        this.name = name;
    }
}
