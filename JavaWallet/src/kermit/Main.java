package kermit;

import javafx.application.Application;
import javafx.fxml.FXMLLoader;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.stage.Stage;

public class Main extends Application {

  @Override
  public void start(Stage primaryStage) throws Exception {
    FXMLLoader loader = new FXMLLoader(getClass().getResource("walletUI.fxml"));
    Parent root = loader.load();
    MainController controller = loader.getController();
    controller.init(primaryStage, "0x1011010010110100", 10000.3f);
    Scene scene = new Scene(root, 700, 500);
    primaryStage.setTitle("Kermit Koin");
    primaryStage.setScene(scene);
    scene.getStylesheets().add(Main.class.getResource("wallet.css").toExternalForm());
    primaryStage.show();
  }

  public static void main(String[] args) {
    launch(args);
  }
}
