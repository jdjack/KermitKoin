package kermit;

import javafx.fxml.FXML;
import javafx.fxml.FXMLLoader;
import javafx.scene.Scene;
import javafx.scene.layout.GridPane;
import javafx.scene.layout.VBox;
import javafx.scene.text.Text;
import javafx.stage.Modality;
import javafx.stage.Stage;

public class MainController {

  Stage primaryStage;
  String address;
  float balance;

  public void init(Stage primaryStage, String address, float balance) {
    this.primaryStage = primaryStage;
    this.address = address;
    this.balance = balance;
    balanceText.setText(String.valueOf(balance));
  }

  @FXML
  private Text balanceText;

  @FXML
  protected void handleSendButtonAction() throws Exception {
    final Stage stage = new Stage();
    stage.initModality(Modality.APPLICATION_MODAL);
    stage.initOwner(primaryStage);
    FXMLLoader loader = new FXMLLoader(getClass().getResource("sendForm.fxml"));
    GridPane sendScreen = loader.load();
    SendController sendController = loader.getController();
    sendController.init(balance, stage);
    Scene sendScene = new Scene(sendScreen, 700, 500);
    stage.setTitle("Send");
    sendScene.getStylesheets().add(Main.class.getResource("wallet.css").toExternalForm());
    stage.setScene(sendScene);
    stage.show();
  }

  @FXML
  protected void handleReceiveButtonAction() throws Exception {
    final Stage stage = new Stage();
    stage.initModality(Modality.APPLICATION_MODAL);
    stage.initOwner(primaryStage);
    FXMLLoader loader = new FXMLLoader(getClass().getResource("receivePopup.fxml"));
    VBox addressScreen = loader.load();
    ReceiveController receiveController = loader.getController();
    Scene addressScene = new Scene(addressScreen, 500, 175);
    stage.setTitle("Receive");
    addressScene.getStylesheets().add(Main.class.getResource("wallet.css").toExternalForm());
    stage.setScene(addressScene);
    receiveController.setAddress(address);
    stage.show();
  }

}
