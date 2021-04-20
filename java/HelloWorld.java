public class HelloWorld{
    public static void main(String[] args){
        System.out.println("Hello World");
        Object o = HelloWorld.class;
        System.out.println(o);
        System.out.println(o.getClass().toString());
    }
}
