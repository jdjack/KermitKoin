package kermit;

import javafx.application.Application;
import javafx.fxml.FXMLLoader;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.stage.Stage;

import java.io.*;
import java.net.MalformedURLException;
import java.net.URL;
import java.net.URLConnection;

public class Main extends Application {

//  public static KeyPair keyPair;
  public static int walletID;
  public static String walletAddress;
  public static float balance;
  public static URL addressURL;
  public static URL balanceURL;

  @Override
  public void start(Stage primaryStage) throws Exception {
    //initMyID();
    //initAddressURL();
    initWalletAddress();
    //initURLs();
    getBalance();
    FXMLLoader loader = new FXMLLoader(getClass().getResource("walletUI.fxml"));
    Parent root = loader.load();
    MainController controller = loader.getController();
    controller.init(primaryStage, walletAddress, balance);
    Scene scene = new Scene(root, 700, 500);
    primaryStage.setTitle("Kermit Koin");
    primaryStage.setScene(scene);
    scene.getStylesheets().add(Main.class.getResource("wallet.css").toExternalForm());
    primaryStage.show();
  }

  public void getBalance() throws IOException {
//    URLConnection connection = addressURL.openConnection();
//    BufferedReader in = new BufferedReader(new InputStreamReader(connection.getInputStream()));
//    balance = Float.parseFloat(in.readLine());
//    System.out.println(balance);
//    in.close();
    balance = 13.37f;
  }

  private void initURLs() throws MalformedURLException {
    balanceURL = new URL("http://129.31.236.46:8082/getBalance?key=" + walletAddress);
  }

  private void initAddressURL() throws MalformedURLException {
    addressURL = new URL("http://129.31.236.46:8082/getAddress?key=" + walletID);
    //addressURL = new URL("http://www.oracle.com/");


  }

  private void initWalletAddress() throws IOException {
//    File file = new File("walletAddress.ser");
//    if (file.createNewFile()) {
//      String myAddress = getWalletAddress();
//      BufferedWriter writer = new BufferedWriter(new FileWriter("walletAddress.ser"));
//      writer.write(myAddress);
//      writer.close();
//    }
//    BufferedReader reader = new BufferedReader(new FileReader("walletAddress.ser"));
//    walletAddress = reader.readLine();
//    reader.close();
    walletAddress = "0x1234567890abcdef";
  }

  private String getWalletAddress() throws IOException {
    URLConnection connection = addressURL.openConnection();
    BufferedReader in = new BufferedReader(new InputStreamReader(connection.getInputStream()));
    String address = in.readLine();
    System.out.println(address);
    in.close();
    return address;
  }

  private void initMyID() throws IOException {
    File file = new File("walletID.ser");
    if (file.createNewFile()) {
      int myID = (int) (System.currentTimeMillis() / 1000);
      String myIDString = String.valueOf(myID);
      BufferedWriter writer = new BufferedWriter(new FileWriter("walletID.ser"));
      writer.write(myIDString);
      writer.close();
    }
    BufferedReader reader = new BufferedReader(new FileReader("walletID.ser"));
    String myIDString = reader.readLine();
    walletID = Integer.parseInt(myIDString);
    reader.close();
  }

//  private void initKeyPair() throws Exception{
//    File file = new File("keys.ser");
//    if (file.createNewFile()) {
//      KeyPair keys = KeyPairGenerator.getInstance("RSA").generateKeyPair();
//      FileOutputStream fileOutputStream = new FileOutputStream("keys.ser");
//      ObjectOutputStream objectOutputStream = new ObjectOutputStream(fileOutputStream);
//      objectOutputStream.writeObject(keys);
//      objectOutputStream.close();
//    }
//    FileInputStream fileInputStream = new FileInputStream("keys.ser");
//    ObjectInputStream objectInputStream = new ObjectInputStream(fileInputStream);
//    keyPair = (KeyPair)objectInputStream.readObject();
//    objectInputStream.close();
//    RSAPublicKey pubK = (RSAPublicKey)(keyPair.getPublic());
//    System.out.println(pubK.getPublicExponent());
//    System.out.println(pubK.getModulus());
//  }

  public static void main(String[] args) {
    launch(args);
  }
}
